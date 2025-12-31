package p2p

import (
	"net"
	"time"
)

func IsPeerAlive(addr string) bool {
	conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
