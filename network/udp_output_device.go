package network

import (
	"github.com/yoojia/go-gecko/v2"
)

func UDPOutputDeviceFactory() (string, gecko.Factory) {
	return "UDPOutputDevice", func() interface{} {
		return NewUDPOutputDevice()
	}
}

func NewUDPOutputDevice() *UDPOutputDevice {
	return &UDPOutputDevice{
		AbcNetworkOutputDevice: NewAbcNetworkOutputDevice("udp"),
	}
}

// UDP客户端输出设备
type UDPOutputDevice struct {
	*AbcNetworkOutputDevice
}
