package fdbased

import (
	"errors"

	"github.com/amnezia-vpn/amnezia-tun2socks/v2/core/device"
)

func Open(name string, mtu uint32) (device.Device, error) {
	return nil, errors.New("not supported")
}
