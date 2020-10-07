package darwin

import "dev.azure.com/bGridSolutions/Tools/ble.git"

func uuidSlice(uu []ble.UUID) [][]byte {
	us := [][]byte{}
	for _, u := range uu {
		us = append(us, ble.Reverse(u))
	}
	return us
}
