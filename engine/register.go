package engine

import (
	_ "github.com/amnezia-vpn/amnezia-tun2socks/v2/proxy/direct"
	_ "github.com/amnezia-vpn/amnezia-tun2socks/v2/proxy/http"
	_ "github.com/amnezia-vpn/amnezia-tun2socks/v2/proxy/reject"
	_ "github.com/amnezia-vpn/amnezia-tun2socks/v2/proxy/relay"
	_ "github.com/amnezia-vpn/amnezia-tun2socks/v2/proxy/shadowsocks"
	_ "github.com/amnezia-vpn/amnezia-tun2socks/v2/proxy/socks4"
	_ "github.com/amnezia-vpn/amnezia-tun2socks/v2/proxy/socks5"
	_ "github.com/amnezia-vpn/amnezia-tun2socks/v2/proxy/ssh"
)
