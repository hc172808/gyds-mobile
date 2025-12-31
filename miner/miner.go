package miner

import (
	"encoding/json"
	"fmt"
	"gyds-mobile/core"
	"io/ioutil"
	"log"
	"os/exec"
	"time"
)

type MinerConfig struct {
	MiningIntervalSeconds int `json:"mining_interval_seconds"`
	MiningDifficulty      int `json:"mining_difficulty"`
}

func LoadConfig(path string) MinerConfig {
	cfg := MinerConfig{
		MiningIntervalSeconds: 5,
		MiningDifficulty:      1,
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("⚠️ Config not found, using defaults")
		return cfg
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		log.Println("⚠️ Invalid config, using defaults")
		return cfg
	}
	return cfg
}

func Start(chain *core.LightChain) {
	fmt.Println("⛏️ Mobile CPU miner started (battery-aware, configurable)")
	cfg := LoadConfig("config.json")

	for {
		batteryLevel := getBatteryLevel()
		if batteryLevel >= 0 && batteryLevel < 20 {
			fmt.Println("🔋 Battery low, pausing mining for 1 min")
			time.Sleep(1 * time.Minute)
			continue
		}

		// Generate a fake block for demo (replace with real PoW)
		latest := chain.GetLatestBlock()
		index := 0
		if latest != nil {
			index = latest.Index + 1
		}

		// Simulate difficulty by sleeping longer for higher difficulty
		time.Sleep(time.Duration(cfg.MiningDifficulty) * time.Second)

		block := &core.Block{
			Index:     index,
			PrevHash:  "",
			Hash:      fmt.Sprintf("blockhash-%d", time.Now().UnixNano()),
			Timestamp: time.Now().Unix(),
			Nonce:     0,
		}
		chain.AddBlock(block)

		time.Sleep(time.Duration(cfg.MiningIntervalSeconds) * time.Second)
	}
}

// getBatteryLevel() remains the same as before
func getBatteryLevel() int {
	out, err := exec.Command("termux-battery-status").Output()
	if err != nil {
		log.Println("⚠️ Unable to read battery status:", err)
		return -1
	}

	var data struct {
		Percentage float64 `json:"percentage"`
		Plugged    bool    `json:"plugged"`
	}
	if err := json.Unmarshal(out, &data); err != nil {
		log.Println("⚠️ Error parsing battery JSON:", err)
		return -1
	}

	if data.Plugged {
		return int(data.Percentage + 100)
	}

	return int(data.Percentage)
}
