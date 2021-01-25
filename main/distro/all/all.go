package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Required features. Can't remove unless there is replacements.
	_ "github.com/nilhost/overnet/app/dispatcher"
	_ "github.com/nilhost/overnet/app/proxyman/inbound"
	_ "github.com/nilhost/overnet/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/nilhost/overnet/app/commander"
	_ "github.com/nilhost/overnet/app/log/command"
	_ "github.com/nilhost/overnet/app/proxyman/command"
	_ "github.com/nilhost/overnet/app/stats/command"

	// Other optional features.
	_ "github.com/nilhost/overnet/app/dns"
	_ "github.com/nilhost/overnet/app/log"
	_ "github.com/nilhost/overnet/app/onet"
	_ "github.com/nilhost/overnet/app/p2p"
	_ "github.com/nilhost/overnet/app/policy"
	_ "github.com/nilhost/overnet/app/reverse"
	_ "github.com/nilhost/overnet/app/router"
	_ "github.com/nilhost/overnet/app/stats"

	// Inbound and outbound proxies.
	_ "github.com/nilhost/overnet/proxy/blackhole"
	_ "github.com/nilhost/overnet/proxy/dns"
	_ "github.com/nilhost/overnet/proxy/dokodemo"
	_ "github.com/nilhost/overnet/proxy/freedom"
	_ "github.com/nilhost/overnet/proxy/http"
	_ "github.com/nilhost/overnet/proxy/mtproto"
	_ "github.com/nilhost/overnet/proxy/shadowsocks"
	_ "github.com/nilhost/overnet/proxy/socks"
	_ "github.com/nilhost/overnet/proxy/trojan"
	_ "github.com/nilhost/overnet/proxy/vless/inbound"
	_ "github.com/nilhost/overnet/proxy/vless/outbound"
	_ "github.com/nilhost/overnet/proxy/vmess/inbound"
	_ "github.com/nilhost/overnet/proxy/vmess/outbound"

	// Transports
	_ "github.com/nilhost/overnet/transport/internet/domainsocket"
	_ "github.com/nilhost/overnet/transport/internet/http"
	_ "github.com/nilhost/overnet/transport/internet/kcp"
	_ "github.com/nilhost/overnet/transport/internet/quic"
	_ "github.com/nilhost/overnet/transport/internet/tcp"
	_ "github.com/nilhost/overnet/transport/internet/tls"
	_ "github.com/nilhost/overnet/transport/internet/udp"
	_ "github.com/nilhost/overnet/transport/internet/websocket"
	_ "github.com/nilhost/overnet/transport/internet/xtls"

	// Transport headers
	_ "github.com/nilhost/overnet/transport/internet/headers/http"
	_ "github.com/nilhost/overnet/transport/internet/headers/noop"
	_ "github.com/nilhost/overnet/transport/internet/headers/srtp"
	_ "github.com/nilhost/overnet/transport/internet/headers/tls"
	_ "github.com/nilhost/overnet/transport/internet/headers/utp"
	_ "github.com/nilhost/overnet/transport/internet/headers/wechat"
	_ "github.com/nilhost/overnet/transport/internet/headers/wireguard"

	// JSON config support. Choose only one from the two below.
	// The following line loads JSON from overctl
	_ "github.com/nilhost/overnet/main/json"
	// The following line loads JSON internally
	// _ "github.com/nilhost/overnet/main/jsonem"

	// Load config from file or http(s)
	_ "github.com/nilhost/overnet/main/confloader/external"
)
