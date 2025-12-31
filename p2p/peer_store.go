package p2p

import "sync"

type PeerStore struct {
	mu    sync.RWMutex
	peers map[string]bool
}

func NewPeerStore() *PeerStore {
	return &PeerStore{
		peers: make(map[string]bool),
	}
}

func (ps *PeerStore) Add(peer string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.peers[peer] = true
}

func (ps *PeerStore) List() []string {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	list := make([]string, 0, len(ps.peers))
	for p := range ps.peers {
		list = append(list, p)
	}
	return list
}
