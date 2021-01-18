// +build !confonly

package quic

import (
	"github.com/nilhost/overnet/common"
	"github.com/nilhost/overnet/transport/internet"
)

//go:generate go run github.com/nilhost/overnet/common/errors/errorgen

// Here is some modification needs to be done before update quic vendor.
// * use bytespool in buffer_pool.go
// * set MaxReceivePacketSize to 1452 - 32 (16 bytes auth, 16 bytes head)
//
//

const protocolName = "quic"
const internalDomain = "quic.internal.v2ray.com"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
