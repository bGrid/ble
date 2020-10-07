package dev

import (
	"dev.azure.com/bGridSolutions/Tools/ble.git"
	"dev.azure.com/bGridSolutions/Tools/ble.git/darwin"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return darwin.NewDevice(opts...)
}
