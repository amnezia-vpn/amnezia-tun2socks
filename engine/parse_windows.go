package engine

import (
	"net/url"

	wun "github.com/amnezia-vpn/amneziawg-go/tun"
	"github.com/gorilla/schema"
	"golang.org/x/sys/windows"

	"github.com/amnezia-vpn/amnezia-tun2socks/v2/core/device"
	"github.com/amnezia-vpn/amnezia-tun2socks/v2/core/device/tun"
	"github.com/amnezia-vpn/amnezia-tun2socks/v2/internal/version"
)

func init() {
	wun.WintunTunnelType = version.Name
}

func parseTUN(u *url.URL, mtu uint32) (device.Device, error) {
	opts := struct {
		GUID string
	}{}
	if err := schema.NewDecoder().Decode(&opts, u.Query()); err != nil {
		return nil, err
	}
	if opts.GUID != "" {
		guid, err := windows.GUIDFromString(opts.GUID)
		if err != nil {
			return nil, err
		}
		wun.WintunStaticRequestedGUID = &guid
	}
	return tun.Open(u.Host, mtu)
}
