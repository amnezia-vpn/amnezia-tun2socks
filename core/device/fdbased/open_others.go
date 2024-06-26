//go:build !linux && !windows

package fdbased

import (
	"fmt"
	"os"

	"github.com/amnezia-vpn/amnezia-tun2socks/v2/core/device"
	"github.com/amnezia-vpn/amnezia-tun2socks/v2/core/device/iobased"
)

func open(fd int, mtu uint32) (device.Device, error) {
	f := &FD{fd: fd, mtu: mtu}

	ep, err := iobased.New(os.NewFile(uintptr(fd), f.Name()), mtu, 0)
	if err != nil {
		return nil, fmt.Errorf("create endpoint: %w", err)
	}
	f.LinkEndpoint = ep

	return f, nil
}
