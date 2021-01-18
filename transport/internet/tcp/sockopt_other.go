// +build !linux,!freebsd
// +build !confonly

package tcp

import (
	"github.com/nilhost/overnet/common/net"
	"github.com/nilhost/overnet/transport/internet"
)

func GetOriginalDestination(conn internet.Connection) (net.Destination, error) {
	return net.Destination{}, nil
}
