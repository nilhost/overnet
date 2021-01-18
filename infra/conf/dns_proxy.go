package conf

import (
	"github.com/golang/protobuf/proto"
	"github.com/nilhost/overnet/common/net"
	"github.com/nilhost/overnet/proxy/dns"
)

type DnsOutboundConfig struct {
	Network Network  `json:"network"`
	Address *Address `json:"address"`
	Port    uint16   `json:"port"`
}

func (c *DnsOutboundConfig) Build() (proto.Message, error) {
	config := &dns.Config{
		Server: &net.Endpoint{
			Network: c.Network.Build(),
			Port:    uint32(c.Port),
		},
	}
	if c.Address != nil {
		config.Server.Address = c.Address.Build()
	}
	return config, nil
}
