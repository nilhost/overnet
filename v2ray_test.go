package core_test

import (
	"testing"

	proto "github.com/golang/protobuf/proto"
	. "github.com/nilhost/overnet"
	"github.com/nilhost/overnet/app/dispatcher"
	"github.com/nilhost/overnet/app/proxyman"
	"github.com/nilhost/overnet/common"
	"github.com/nilhost/overnet/common/net"
	"github.com/nilhost/overnet/common/protocol"
	"github.com/nilhost/overnet/common/serial"
	"github.com/nilhost/overnet/common/uuid"
	"github.com/nilhost/overnet/features/dns"
	"github.com/nilhost/overnet/features/dns/localdns"
	_ "github.com/nilhost/overnet/main/distro/all"
	"github.com/nilhost/overnet/proxy/dokodemo"
	"github.com/nilhost/overnet/proxy/vmess"
	"github.com/nilhost/overnet/proxy/vmess/outbound"
	"github.com/nilhost/overnet/testing/servers/tcp"
)

func TestV2RayDependency(t *testing.T) {
	instance := new(Instance)

	wait := make(chan bool, 1)
	instance.RequireFeatures(func(d dns.Client) {
		if d == nil {
			t.Error("expected dns client fulfilled, but actually nil")
		}
		wait <- true
	})
	instance.AddFeature(localdns.New())
	<-wait
}

func TestV2RayClose(t *testing.T) {
	port := tcp.PickPort()

	userId := uuid.New()
	config := &Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
		Inbound: []*InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortRange: net.SinglePortRange(port),
					Listen:    net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address: net.NewIPOrDomain(net.LocalHostIP),
					Port:    uint32(0),
					NetworkList: &net.NetworkList{
						Network: []net.Network{net.Network_TCP, net.Network_UDP},
					},
				}),
			},
		},
		Outbound: []*OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&outbound.Config{
					Receiver: []*protocol.ServerEndpoint{
						{
							Address: net.NewIPOrDomain(net.LocalHostIP),
							Port:    uint32(0),
							User: []*protocol.User{
								{
									Account: serial.ToTypedMessage(&vmess.Account{
										Id: userId.String(),
									}),
								},
							},
						},
					},
				}),
			},
		},
	}

	cfgBytes, err := proto.Marshal(config)
	common.Must(err)

	server, err := StartInstance("protobuf", cfgBytes)
	common.Must(err)
	server.Close()
}
