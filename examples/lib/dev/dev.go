package dev

import (
	"dev.azure.com/bGridSolutions/Tools/ble.git"
)

// NewDevice ...
func NewDevice(impl string, opts ...ble.Option) (d ble.Device, err error) {
	return DefaultDevice(opts...)
}
