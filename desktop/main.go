package main

import (
	"encoding/base64"
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

	"gx/ipfs/QmUJYo4etAQqFfSS2rarFAE97eNGB8ej64YkRT2SmsYD4r/go-ipfs/repo/fsrepo"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
	"github.com/mitchellh/go-homedir"
	"github.com/skip2/go-qrcode"
	"github.com/textileio/textile-go/core"
	"github.com/textileio/textile-go/gateway"
	"github.com/textileio/textile-go/ipfs"
	"github.com/textileio/textile-go/keypair"
	"github.com/textileio/textile-go/repo"
)

var (
	appName     = "Textile"
	builtAt     string
	debug       = flag.Bool("d", false, "enables the debug mode")
	window      *astilectron.Window
	menu        *astilectron.Menu
	gatewayAddr string
	repoPath    string
	expanded    bool
	onboarded   = false
	connected   = false
	config      *Config
)

const (
	SetupSize       = 384
	QRCodeSize      = 256
	InitialWidth    = 1024
	InitialHeight   = 633
	SleepOnLoad     = time.Second * 1
	SleepOnPreReady = time.Millisecond * 200
	SleepOnExpand   = time.Millisecond * 200
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
	repoPath = filepath.Join(appDir, "repo")

	// ignore errors and assume defaults
	if config, err = ReadConfig(repoPath); err != nil {
		config = &Config{}
	}

	sendData("status", map[string]interface{}{
					"status": "loading",
				})

	// run init if needed
	if !fsrepo.IsInitialized(repoPath) {
		accnt := keypair.Random()
		initc := core.InitConfig{
			Account:   accnt,
			RepoPath:  repoPath,
			LogToDisk: true,
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

	gateway.Host = &gateway.Gateway{
		Node: node,
	}

	// bring the node online and startup the gateway
	if err := node.Start(); err != nil {
		return err
	}
	<-node.OnlineCh()

	astilog.Info(fmt.Sprintf("Peer ID: %s", node.Ipfs().Identity.Pretty()))

	menu = t.NewMenu([]*astilectron.MenuItemOptions{
		{
			Label:   astilectron.PtrStr("Backup"),
			Enabled: astilectron.PtrBool(false),
		},
		{
			Label: astilectron.PtrStr("Choose location..."),
			OnClick: func(e astilectron.Event) (deleteListener bool) {
				astilog.Info("Updating Backup Location...")
				sendMessage("set-backup")
				window.Focus()
				return
			},
		},
		{
			Label: astilectron.PtrStr("Open..."),
			OnClick: func(e astilectron.Event) (deleteListener bool) {
				astilog.Info("Opening backup folder...")
				browser.OpenFile(config.BackupFolder)
				return
			},
		},
		{
			Type: astilectron.MenuItemTypeSeparator,
		},
		{
			Label: astilectron.PtrStr("Check Messages"),
			OnClick: func(e astilectron.Event) (deleteListener bool) {
				astilog.Info("Checking Messages...")
				node.CheckCafeMessages()
				return
			},
		},
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
				window.Log("before")
				sendData("thread.update", map[string]interface{}{
					"update": mapped,
				})
				window.Log("after")
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
				username := node.ContactUsername(note.ActorId)
				var uinote = a.NewNotification(&astilectron.NotificationOptions{
					Title: note.Subject,
					Body:  fmt.Sprintf("%s: %s.", username, note.Body),
					Icon:  "/resources/icon.png",
				})

				// tmp auto-accept thread invites
				if note.Type == repo.InviteReceivedNotification.Description() {
					go func(tid string) {
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

	// start the gateway
	gateway.Host.Start(fmt.Sprintf("127.0.0.1:%s", core.GetRandomPort()))

	// save off the server address
	gatewayAddr = fmt.Sprintf("http://%s", gateway.Host.Addr())
	astilog.Info(fmt.Sprintf("Gateway Addr: %s", gatewayAddr))

	// sleep for a bit on the landing screen, it feels better
	time.Sleep(SleepOnLoad)

	sendData("status", map[string]interface{}{
					"status": "ready",
				})

	// check if we're configured yet
	threads := node.Threads()
	if len(threads) > 0 {
		astilog.Info("Have thread(s):")
		for _, thread := range threads {
			astilog.Info(thread.Id)
		}
	} else {
		astilog.Info("No thread(s) yet")
		// get qr code for setup
		qr, pk, err := getQRCode()
		if err != nil {
			astilog.Error(err)
			return err
		}
		sendData("pair", map[string]interface{}{
			"qr": qr,
			"pk": pk,
		})
	}

	return nil
}

func sendMessage(name string) {
	window.SendMessage(map[string]string{"name": name})
}

func sendData(name string, data map[string]interface{}) {
	data["name"] = name
	window.SendMessage(data)
}

type ThreadPayload struct {
	ThreadId string `json:"thread_id"`
	Body string `json:"body"`
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
	case "threads":
		// TODO: Use API (textile threads ls)
		list := make([]*core.ThreadInfo, 0)
		for _, thrd := range node.Threads() {
			info, err := thrd.Info()
			if err != nil {
				return nil, err
			}
			list = append(list, info)
		}
		mapped, err := toMap(map[string]interface{}{
			"threads": list,
		})
		if err != nil {
			return nil, err
		}
		return mapped, nil
	case "backup":
		var folder string
		if err := json.Unmarshal(m.Payload, &folder); err != nil {
			return nil, err
		}
		config.BackupFolder = folder
		WriteConfig(repoPath, config)
		return map[string]interface{}{
			"success": true,
		}, nil
	case "status":
		status := "loading"
		if node != nil {
			status = "ready"
		}
		return map[string]interface{}{
			"status": status,
		}, nil
	case "username":
		// TODO: Use API (textile profile get)
		var username string
		name, err := node.Username()
		if err == nil && name != nil && *name != "" {
			username = *name
		} else {
			id := node.Ipfs().Identity.Pretty()
			username = ipfs.ShortenID(id)
		}
		return map[string]interface{}{
			"username": username,
		}, nil
	case "messages":
		// TODO: Use API (textile messages ls -t -o -l)
		var threadId string
		if err := json.Unmarshal(m.Payload, &threadId); err != nil {
			return nil, err
		}
		list, err := node.ThreadMessages("", 20, threadId)
		if err != nil {
			return nil, err
		}
		mapped, err := toMap(map[string]interface{}{
			"messages": list,
		})
		if err != nil {
			return nil, err
		}
		return mapped, nil
	case "message":
		var payload ThreadPayload
		if err := json.Unmarshal(m.Payload, &payload); err != nil {
			return nil, err
		}
		if payload.ThreadId == "default" {
			payload.ThreadId = node.Config().Threads.Defaults.ID
		}
		thrd := node.Thread(payload.ThreadId)
		if thrd == nil {
			return nil, core.ErrThreadNotFound
		}
		hash, err := thrd.AddMessage(payload.Body)
		if err != nil {
			return nil, err
		}
		info, err := node.BlockInfo(hash.B58String())
		if err != nil {
			return nil, err
		}
		mapped, err := toMap(info)
		if err != nil {
			return nil, err
		}
		return mapped, nil
	default:
		return map[string]interface{}{}, nil
	}
}

func getQRCode() (string, string, error) {
	// get our own peer id for receiving an account key
	pid, err := node.PeerId()
	if err != nil {
		return "", "", err
	}

	// create a qr code
	url := fmt.Sprintf("https://www.textile.photos/invites/device#id=%s", pid.Pretty())
	png, err := qrcode.Encode(url, qrcode.Medium, QRCodeSize)
	if err != nil {
		return "", "", err
	}

	return base64.StdEncoding.EncodeToString(png), pid.Pretty(), nil
}

func sendPreReady() {
	sendMessage("preready")
	time.Sleep(SleepOnPreReady)
}

func Exit(e astilectron.Event) (deleteListener bool) {
	astilog.Info("Exit clicked")
	window.Close()
	return
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
