package tun

import (
	"github.com/amnezia-vpn/amneziawg-go/tun"
)

const (
	offset     = 0
	defaultMTU = 0 /* auto */
)

func createTUN(name string, mtu int) (tun.Device, error) {
	return tun.CreateTUN(name, mtu)
}
