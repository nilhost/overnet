// +build !confonly

package tcp

import (
	"github.com/nilhost/overnet/common"
	"github.com/nilhost/overnet/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
