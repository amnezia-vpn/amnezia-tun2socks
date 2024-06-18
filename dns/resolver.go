package dns

import (
	"net"

	"github.com/amnezia-vpn/amnezia-tun2socks/v2/dialer"
)

func init() {
	// We must use this DialContext to query DNS
	// when using net default resolver.
	net.DefaultResolver.PreferGo = true
	net.DefaultResolver.Dial = dialer.DefaultDialer.DialContext
}
