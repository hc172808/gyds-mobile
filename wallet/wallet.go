package wallet

import (
	"crypto/ed25519"
	"encoding/hex"
)

type Wallet struct {
	PrivateKey ed25519.PrivateKey
	PublicKey  ed25519.PublicKey
}

func NewWallet() *Wallet {
	pub, priv, _ := ed25519.GenerateKey(nil)
	return &Wallet{
		PrivateKey: priv,
		PublicKey:  pub,
	}
}

func (w *Wallet) Address() string {
	return hex.EncodeToString(w.PublicKey)
}

func (w *Wallet) Sign(data []byte) []byte {
	return ed25519.Sign(w.PrivateKey, data)
}

func (w *Wallet) Verify(sig, data []byte) bool {
	return ed25519.Verify(w.PublicKey, data, sig)
}
