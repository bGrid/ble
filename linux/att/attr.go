package att

import "dev.azure.com/bGridSolutions/Tools/ble.git"

// attr is a BLE attribute.
type attr struct {
	h    uint16
	endh uint16
	typ  ble.UUID

	v  []byte
	rh ble.ReadHandler
	wh ble.WriteHandler
}
