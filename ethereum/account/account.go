/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2021-01-25
 */
package account

import "github.com/nilhost/overnet/ethereum/account/wallet"

func NewWallet(path string, password string) error {
	return wallet.GenerateAndSaveEthereumWallet(path, password)
}