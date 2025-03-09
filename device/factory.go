package device

import (
	"fmt"
)

func NewDevice(name string) (*Device, error) {
	dt := DeviceType(name)
	device, exist := DeviceMap[dt]

	if !exist {
		return &device, fmt.Errorf("unrecognised device %v", name)
	}

	return &device, nil
}

func DefaultDevice() *Device {
	device := DeviceMap[DEFAULT]
	return &device
}
