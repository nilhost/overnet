/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2019-09-06
 */
package protocol

import "github.com/nilhost/overnet/app/p2p/protocol/seedlist"

type Protocol struct {
	*seedlist.SeedListProtocol
}
