package mobile

import (
	"crypto/rand"
	libp2pc "gx/ipfs/QmTW4SdgBWq9GjsBsHeUx8WuGxzhgzAf88UMH2w62PC8yK/go-libp2p-crypto"

	"github.com/golang/protobuf/proto"
	"github.com/textileio/go-textile/core"
	"github.com/textileio/go-textile/pb"
)

// AddThread adds a new thread with the given name
func (m *Mobile) AddThread(config []byte) ([]byte, error) {
	if !m.node.Started() {
		return nil, core.ErrStopped
	}

	conf := new(pb.AddThreadConfig)
	if err := proto.Unmarshal(config, conf); err != nil {
		return nil, err
	}

	sk, _, err := libp2pc.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, err
	}

	thrd, err := m.node.AddThread(*conf, sk, m.node.Account().Address(), true, true)
	if err != nil {
		return nil, err
	}

	view, err := m.node.ThreadView(thrd.Id)
	if err != nil {
		return nil, err
	}

	return proto.Marshal(view)
}

// AddOrUpdateThread calls core AddOrUpdateThread
func (m *Mobile) AddOrUpdateThread(thrd []byte) error {
	if !m.node.Online() {
		return core.ErrOffline
	}

	mthrd := new(pb.Thread)
	if err := proto.Unmarshal(thrd, mthrd); err != nil {
		return err
	}

	return m.node.AddOrUpdateThread(mthrd)
}

// RenameThread call core RenameThread
func (m *Mobile) RenameThread(id string, name string) error {
	if !m.node.Started() {
		return core.ErrStopped
	}

	return m.node.RenameThread(id, name)
}

// Thread calls core Thread
func (m *Mobile) Thread(id string) ([]byte, error) {
	if !m.node.Started() {
		return nil, core.ErrStopped
	}

	view, err := m.node.ThreadView(id)
	if err != nil {
		return nil, err
	}
	return proto.Marshal(view)
}

// Threads lists all threads
func (m *Mobile) Threads() ([]byte, error) {
	if !m.node.Started() {
		return nil, core.ErrStopped
	}

	views := &pb.ThreadList{
		Items: make([]*pb.Thread, 0),
	}
	for _, thrd := range m.node.Threads() {
		view, err := m.node.ThreadView(thrd.Id)
		if err != nil {
			return nil, err
		}
		views.Items = append(views.Items, view)
	}

	return proto.Marshal(views)
}

// ThreadPeers calls core ThreadPeers
func (m *Mobile) ThreadPeers(id string) ([]byte, error) {
	if !m.node.Started() {
		return nil, core.ErrStopped
	}

	peers, err := m.node.ThreadPeers(id)
	if err != nil {
		return nil, err
	}

	return proto.Marshal(peers)
}

// RemoveThread call core RemoveThread
func (m *Mobile) RemoveThread(id string) (string, error) {
	if !m.node.Started() {
		return "", core.ErrStopped
	}

	hash, err := m.node.RemoveThread(id)
	if err != nil {
		return "", err
	}

	return hash.B58String(), nil
}

// SnapshotThreads calls core SnapshotThreads
func (m *Mobile) SnapshotThreads() error {
	if !m.node.Started() {
		return core.ErrStopped
	}

	return m.node.SnapshotThreads()
}

// SearchThreadSnapshots calls core SearchThreadSnapshots
func (m *Mobile) SearchThreadSnapshots(query []byte, options []byte) (*SearchHandle, error) {
	if !m.node.Online() {
		return nil, core.ErrOffline
	}

	mquery := new(pb.ThreadSnapshotQuery)
	if err := proto.Unmarshal(query, mquery); err != nil {
		return nil, err
	}
	moptions := new(pb.QueryOptions)
	if err := proto.Unmarshal(options, moptions); err != nil {
		return nil, err
	}

	resCh, errCh, cancel, err := m.node.SearchThreadSnapshots(mquery, moptions)
	if err != nil {
		return nil, err
	}

	return m.handleSearchStream(resCh, errCh, cancel)
}
