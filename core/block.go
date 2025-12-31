package core

import "time"

type Block struct {
	Index        int
	PrevHash     string
	Hash         string
	Timestamp    int64
	MinerAddress string
	Nonce        int
}
