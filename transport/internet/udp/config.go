package udp

import (
	"github.com/nilhost/overnet/common"
	"github.com/nilhost/overnet/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
