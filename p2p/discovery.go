package p2p

import (
	"log"
	"time"
)

type Discovery struct {
	store *PeerStore
}

func NewDiscovery() *Discovery {
	ps := NewPeerStore()
	for _, p := range BootstrapPeers {
		ps.Add(p)
	}
	return &Discovery{store: ps}
}

func (d *Discovery) Start() {
	go func() {
		for {
			for _, peer := range d.store.List() {
				if IsPeerAlive(peer) {
					log.Println("✅ Peer alive:", peer)
				}
			}
			time.Sleep(10 * time.Second)
		}
	}()
}

func (d *Discovery) Peers() []string {
	return d.store.List()
}
