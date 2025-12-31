package rpc

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Client struct {
	peers []string
}

func NewClient(peers []string) *Client {
	return &Client{peers: peers}
}

func (c *Client) GetBalance(address string) float64 {
	if len(c.peers) == 0 {
		return 0
	}
	url := "https://" + c.peers[0] + "/getBalance?address=" + address
	resp, err := http.Get(url)
	if err != nil {
		log.Println("⚠️ RPC error:", err)
		return 0
	}
	defer resp.Body.Close()
	var data struct {
		Balance float64 `json:"balance"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Println("⚠️ Decode error:", err)
		return 0
	}
	return data.Balance
}

func (c *Client) SendTx(payload map[string]interface{}) bool {
	if len(c.peers) == 0 {
		return false
	}
	url := "https://" + c.peers[0] + "/sendTx"
	data, _ := json.Marshal(payload)
	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		log.Println("⚠️ RPC error:", err)
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == 200
}
