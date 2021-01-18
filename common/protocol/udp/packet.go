package udp

import (
	"github.com/nilhost/overnet/common/buf"
	"github.com/nilhost/overnet/common/net"
)

// Packet is a UDP packet together with its source and destination address.
type Packet struct {
	Payload *buf.Buffer
	Source  net.Destination
	Target  net.Destination
}
