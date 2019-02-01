package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/mr-tron/base58/base58"
	"github.com/pkg/browser"
	"github.com/textileio/textile-go/crypto"

	"gx/ipfs/QmX9YciaxRii8TARoEbmavzaeTUAe7BozeAgydsThNcTpy/go-ipfs/repo/fsrepo"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
	"github.com/mitchellh/go-homedir"
	"github.com/textileio/textile-go/core"
	"github.com/textileio/textile-go/cmd"
	"github.com/textileio/textile-go/gateway"
	"github.com/textileio/textile-go/ipfs"
	"github.com/textileio/textile-go/keypair"
	"github.com/textileio/textile-go/repo"
)

var (
	appName     = "Textile"
	debug       = flag.Bool("d", false, "enables the debug mode")
	window      *astilectron.Window
	menu        *astilectron.Menu
	config      *Config
)

const (
	SetupSize   = 384
	SleepOnLoad = time.Second * 1
)

var node *core.Textile

func main() {
	flag.Parse()
	astilog.FlagInit()
	bootstrapApp()
}

func start(a *astilectron.Astilectron, w []*astilectron.Window, _ *astilectron.Menu, t *astilectron.Tray, _ *astilectron.Menu) error {
	window = w[0]
	window.Show()
	sendData("status", map[string]interface{}{"status": "initializing"})

	// get homedir
	home, err := homedir.Dir()
	if err != nil {
		astilog.Fatal(fmt.Errorf("get homedir failed: %s", err))
	}

	// ensure app support folder is created
	var appDir string
	if runtime.GOOS == "darwin" {
		appDir = filepath.Join(home, "Library/Application Support/Textile")
	} else {
		appDir = filepath.Join(home, ".textile")
	}
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return err
	}
	repoPath := filepath.Join(appDir, "repo")

	// ignore errors and assume defaults
	if config, err = ReadConfig(repoPath); err != nil {
		config = &Config{}
	}

	// run init if needed
	if !fsrepo.IsInitialized(repoPath) {
		accnt := keypair.Random()
		initc := core.InitConfig{
			Account:   accnt,
			RepoPath:  repoPath,
			LogToDisk: true,
			GatewayAddr: fmt.Sprintf("127.0.0.1:%s", core.GetRandomPort()),
			ApiAddr: fmt.Sprintf("127.0.0.1:%s", core.GetRandomPort()),
		}
		if err := core.InitRepo(initc); err != nil {
			return err
		}
	}

	// build textile node
	node, err = core.NewTextile(core.RunConfig{RepoPath: repoPath})
	if err != nil {
		return err
	}

	sendData("status", map[string]interface{}{"status": "starting"})

	// bring the node online
	if err := node.Start(); err != nil {
		return err
	}

	// setup and start the apis
	gateway.Host = &gateway.Gateway{
		Node: node,
	}
	node.StartApi(node.Config().Addresses.API)
	gateway.Host.Start(node.Config().Addresses.Gateway)
	
	<-node.OnlineCh()

	menu = t.NewMenu([]*astilectron.MenuItemOptions{
		{
			Label: astilectron.PtrStr("Quit"),
			OnClick: func(e astilectron.Event) (deleteListener bool) {
				astilog.Info("Quitting...")
				a.Quit()
				return
			},
		},
	})
	menu.Create()

	// subscribe to thread updates
	listener := node.ThreadUpdateListener()
	go func() {
		for {
			select {
			case update, ok := <-listener.Ch:
				if !ok {
					return
				}
				mapped, err := toMap(update)
				if err != nil {
					continue
				}
				sendData("thread.update", map[string]interface{}{
					"update": mapped,
				})
				if config.BackupFolder != "" {
					if data, ok := update.(core.ThreadUpdate); ok {
						if data.Block.Type == "FILES" {
							tmpBase := filepath.Join(config.BackupFolder, data.Block.Target)
							err := backupFile(data.Block.Id, node, tmpBase)
							if err != nil {
								continue
							}
						}
					}
				}
			}
		}
	}()

	// subscribe to notifications
	go func() {
		for {
			select {
			case note, ok := <-node.NotificationCh():
				if !ok {
					return
				}
				username, avatar := node.ContactDisplayInfo(note.ActorId)
				var uinote = a.NewNotification(&astilectron.NotificationOptions{
					Title: note.Subject,
					Body:  fmt.Sprintf("%s: %s.", username, note.Body),
					Icon:  avatar,
				})

				// tmp auto-accept thread invites
				if note.Type == repo.InviteReceivedNotification.Description() {
					go func(tid string) {
						astilog.Info(tid)
						if _, err := node.AcceptThreadInvite(tid); err != nil {
							astilog.Error(err)
						}
					}(note.BlockId)
				}

				// show notification
				go func(n *astilectron.Notification) {
					if err := n.Create(); err != nil {
						astilog.Error(err)
						return
					}
					if err := n.Show(); err != nil {
						astilog.Error(err)
						return
					}
				}(uinote)
			}
		}
	}()

	// print information to terminal
	printSplash()

	// sleep for a bit on the landing screen, it feels better
	time.Sleep(SleepOnLoad)

	// all ready on the js side now!
	sendData("status", map[string]interface{}{"status": "ready"})

	// check if we're configured yet
	threads := node.Threads()
	if len(threads) > 0 {
		astilog.Info("Have thread(s):")
		for _, thread := range threads {
			astilog.Info(thread.Id)
		}
	} else {
		astilog.Info("No thread(s) yet")
	}

	return nil
}

func sendData(name string, data map[string]interface{}) {
	data["name"] = name
	window.SendMessage(data)
}

func toMap(data interface{}) (map[string]interface{}, error) {
	var target map[string]interface{}
	inrec, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &target)
	if err != nil {
		return nil, err
	}
	return target, nil
}

func handleMessage(_ *astilectron.Window, m bootstrap.MessageIn) (interface{}, error) {
	astilog.Info(m.Name)
	switch m.Name {
	case "setBackupFolder":
		var folder string
		if err := json.Unmarshal(m.Payload, &folder); err != nil {
			return nil, err
		}
		config.BackupFolder = folder
		WriteConfig(node.RepoPath(), config)
		return true, nil
	case "openBackupFolder":
		browser.OpenFile(config.BackupFolder)
		return true, nil
	default:
		return map[string]interface{}{}, nil
	}
	return map[string]interface{}{}, nil
}

func backupFile(id string, node *core.Textile, tmpBase string) error {
	info, err := node.ThreadFile(id)
	if err != nil {
		astilog.Error(err)
		return err
	}
	if err = os.MkdirAll(tmpBase, 0744); err != nil {
		astilog.Error(err)
		return err
	}
	for _, file := range info.Files {
		err = os.Mkdir(filepath.Join(tmpBase, strconv.Itoa(file.Index)), 0744)
		if err != nil {
			astilog.Error(err)
			return err
		}
		for name, link := range file.Links {
			data, err := ipfs.DataAtPath(node.Ipfs(), link.Hash)
			if err != nil {
				astilog.Error(err)
				continue
			}
			var plaintext []byte
			if link.Key != "" {
				key, err := base58.Decode(link.Key)
				if err != nil {
					astilog.Error(err)
					continue
				}
				plaintext, err = crypto.DecryptAES(data, key)
			} else {
				plaintext = data
			}
			// TODO: Extract extension from Media value
			ext := strings.ToLower(filepath.Ext(link.Name))
			tempPath := filepath.Join(tmpBase, strconv.Itoa(file.Index), name+ext)
			err = ioutil.WriteFile(tempPath, plaintext, 0644)
			if err != nil {
				astilog.Error(err)
			}
		}
	}
	return nil
}

func printSplash() {
	pid, err := node.PeerId()
	if err != nil {
		astilog.Fatalf("get peer id failed: %s", err)
	}
	fmt.Println(cmd.Grey("Textile desktop version v0.0.1"))
	fmt.Println(cmd.Grey("Repo:    ") + cmd.Grey(node.RepoPath()))
	fmt.Println(cmd.Grey("API:     ") + cmd.Grey(node.ApiAddr()))
	fmt.Println(cmd.Grey("Gateway: ") + cmd.Grey(gateway.Host.Addr()))
	fmt.Println(cmd.Grey("PeerID:  ") + cmd.Green(pid.Pretty()))
	fmt.Println(cmd.Grey("Account: ") + cmd.Cyan(node.Account().Address()))
}
