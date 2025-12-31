package core

import (
	"log"
	"sync"
)

type LightChain struct {
	mu     sync.RWMutex
	blocks []*Block
	maxLen int
}

func NewLightChain() *LightChain {
	return &LightChain{
		blocks: make([]*Block, 0),
		maxLen: 100, // store last 100 blocks only
	}
}

func (lc *LightChain) Start() {
	log.Println("🧱 Lightchain started")
}

func (lc *LightChain) AddBlock(b *Block) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.blocks = append(lc.blocks, b)
	if len(lc.blocks) > lc.maxLen {
		lc.blocks = lc.blocks[1:] // prune old blocks
	}
}

func (lc *LightChain) GetLatestBlock() *Block {
	lc.mu.RLock()
	defer lc.mu.RUnlock()
	if len(lc.blocks) == 0 {
		return nil
	}
	return lc.blocks[len(lc.blocks)-1]
}
