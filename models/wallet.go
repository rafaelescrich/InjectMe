package models
import (
	"InjectMe/utils"
)

type Wallet struct {
	PublicKey string `json:"public_key"`
	User User `json:"user"`
	Balance float32 `json:"balance"`
	UpdatedAt string `json:"updated_at"`
}

func (wallet  *Wallet) GeneratePublicKey(){
	wallet.PublicKey = utils.Md5(User.Username + User.Password)
}
