package miner

import (
	"fmt"
	"gyds-mobile/core"
	"time"
)

func Start(chain *core.LightChain) {
	fmt.Println("⛏️ Mobile CPU miner started")
	for {
		// Generate a fake block for demo (replace with real PoW)
		latest := chain.GetLatestBlock()
		index := 0
		if latest != nil {
			index = latest.Index + 1
		}

		block := &core.Block{
			Index:     index,
			PrevHash:  "",
			Hash:      fmt.Sprintf("blockhash-%d", time.Now().UnixNano()),
			Timestamp: time.Now().Unix(),
			Nonce:     0,
		}
		chain.AddBlock(block)

		time.Sleep(5 * time.Second) // adjustable mining interval
	}
}
