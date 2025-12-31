package main

import (
	"log"
	"time"

	"gyds-mobile/core"
	"gyds-mobile/p2p"
	"gyds-mobile/rpc"
	"gyds-mobile/miner"
)

func main() {
	log.Println("🚀 GYDS Mobile Lite Node Starting")

	// Start P2P discovery to fullnodes
	discovery := p2p.NewDiscovery()
	discovery.Start()

	// Start lightchain
	chain := core.NewLightChain()
	chain.Start()

	// Start CPU-friendly miner
	go miner.Start(chain)

	// Start RPC client to fullnodes
	client := rpc.NewClient(discovery.Peers())

	// Example loop: print balance every 30s
	for {
		bal := client.GetBalance("YOUR_ADDRESS_HERE")
		log.Println("💰 Balance:", bal)
		time.Sleep(30 * time.Second)
	}
}
