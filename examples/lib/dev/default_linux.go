package dev

import (
	"dev.azure.com/bGridSolutions/Tools/ble.git"
	"dev.azure.com/bGridSolutions/Tools/ble.git/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}
