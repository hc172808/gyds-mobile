package miner

import (
	"encoding/json"
	"fmt"
	"gyds-mobile/core"
	"log"
	"os/exec"
	"strconv"
	"time"
)

func Start(chain *core.LightChain) {
	fmt.Println("⛏️ Mobile CPU miner started (battery-aware)")

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

// getBatteryLevel returns battery percentage or -1 if unavailable
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

	// If charging, allow mining even if battery < 20%
	if data.Plugged {
		return int(data.Percentage + 100) // return >20 to allow mining
	}

	return int(data.Percentage)
}
