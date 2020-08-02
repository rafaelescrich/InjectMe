package models

import (
	utils "github.com/OCFloripa/InjectMe/utils"
)

// Wallet information
type Wallet struct {
	PublicKey string  `json:"public_key"`
	User      User    `json:"user"`
	Balance   float32 `json:"balance"`
	UpdatedAt string  `json:"updated_at"`
}

// GeneratePubKey generates the public key
func (wallet *Wallet) GeneratePubKey() {
	wallet.PublicKey = utils.Hash(wallet.User.Username + wallet.User.Password)
}
