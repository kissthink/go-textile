package core

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/textileio/go-textile/broadcast"
	"github.com/textileio/go-textile/crypto"
	"github.com/textileio/go-textile/keypair"
	"github.com/textileio/go-textile/pb"
)

// Account returns account keypair
func (t *Textile) Account() *keypair.Full {
	return t.account
}

// Sign signs input with account seed
func (t *Textile) Sign(input []byte) ([]byte, error) {
	return t.account.Sign(input)
}

// Verify verifies input with account address
func (t *Textile) Verify(input []byte, sig []byte) error {
	return t.account.Verify(input, sig)
}

// Encrypt encrypts input with account address
func (t *Textile) Encrypt(input []byte) ([]byte, error) {
	pk, err := t.account.LibP2PPubKey()
	if err != nil {
		return nil, err
	}
	return crypto.Encrypt(pk, input)
}

// Decrypt decrypts input with account address
func (t *Textile) Decrypt(input []byte) ([]byte, error) {
	sk, err := t.account.LibP2PPrivKey()
	if err != nil {
		return nil, err
	}
	return crypto.Decrypt(sk, input)
}

// AccountThread returns the account private thread
func (t *Textile) AccountThread() *Thread {
	return t.ThreadByKey(t.config.Account.Address)
}

// AccountContact returns a contact for this account
func (t *Textile) AccountContact() *pb.Contact {
	return t.contact(t.account.Address(), false)
}

// SyncAccount performs a thread backup search and applies the result
func (t *Textile) SyncAccount() (*broadcast.Broadcaster, error) {
	query := &pb.ThreadSnapshotQuery{
		Address: t.account.Address(),
	}
	options := &pb.QueryOptions{
		Limit: -1,
		Wait:  5,
	}

	resCh, errCh, cancel, err := t.SearchThreadSnapshots(query, options)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case res, ok := <-resCh:
				if !ok {
					return
				}
				if err := t.applySnapshot(res); err != nil {
					log.Errorf("error applying snap %s: %s", res.Id, err)
				}

			case err := <-errCh:
				log.Errorf("error during account sync: %s", err)
			}
		}
	}()

	return cancel, err
}

// maybeSyncAccount runs SyncAccount if it has not been run in the last kSyncAccountFreq
func (t *Textile) maybeSyncAccount() {
	if t.cancelSync != nil {
		t.cancelSync.Close()
		t.cancelSync = nil
	}

	daily, err := t.datastore.Config().GetLastDaily()
	if err != nil {
		log.Errorf("error getting last daily run: %s", err)
		return
	}

	if daily.Add(kSyncAccountFreq).Before(time.Now()) {
		var err error
		t.cancelSync, err = t.SyncAccount()
		if err != nil {
			log.Errorf("error running sync account: %s", err)
			return
		}

		if err := t.datastore.Config().SetLastDaily(); err != nil {
			log.Errorf("error setting last daily run: %s", err)
		}
	}
}

// accountPeers returns all known account peers
func (t *Textile) accountPeers() []*pb.Peer {
	query := fmt.Sprintf("address='%s' and id!='%s'", t.account.Address(), t.node.Identity.Pretty())
	return t.datastore.Peers().List(query)
}

// applySnapshot unmarshals and adds an unencrypted thread snapshot from a search result
func (t *Textile) applySnapshot(result *pb.QueryResult) error {
	snap := new(pb.Thread)
	if err := ptypes.UnmarshalAny(result.Value, snap); err != nil {
		return err
	}

	log.Debugf("applying snapshot %s", snap.Id)

	return t.AddOrUpdateThread(snap)
}
