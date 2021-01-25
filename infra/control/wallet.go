/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2021-01-25
 */


package control

import (
	"github.com/nilhost/overnet/common"
	"github.com/nilhost/overnet/ethereum/account"
)

type WalletCommand struct{}

func (c *WalletCommand) Name() string {
	return "wallet"
}

func (c *WalletCommand) Description() Description {
	return Description{
		Short: "Generate new wallet",
		Usage: []string{"overctl wallet"},
	}
}

func (c *WalletCommand) Execute(args []string) error {
	if len(args)!=2 {
		panic("args must be two.")
	}
	return account.NewWallet(args[0], args[1])
}

func init() {
	common.Must(RegisterCommand(&WalletCommand{}))
}
