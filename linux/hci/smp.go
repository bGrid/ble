package hci

import (
	"time"

	"dev.azure.com/bGridSolutions/Tools/ble.git"
)

type smpDispatcher struct {
	desc    string
	handler func(*Conn, pdu) ([]byte, error)
}

const (
	IoCapsDisplayOnly     = 0x00
	IoCapsDisplayYesNo    = 0x01
	IoCapsKeyboardOnly    = 0x02
	IoCapsNone            = 0x03
	IoCapsKeyboardDisplay = 0x04
	IoCapsReservedStart   = 0x05
)

type OobDataFlag byte

const (
	OobNotPresent OobDataFlag = iota
	OobPreset
)

type SmpManagerFactory interface {
	Create(SmpConfig) SmpManager
	SetBondManager(BondManager)
}

type SmpManager interface {
	InitContext(localAddr, remoteAddr []byte,
		localAddrType, remoteAddrType uint8)
	Handle(data []byte) error
	Pair(authData ble.AuthData, to time.Duration) error
	BondInfoFor(addr string) BondInfo
	DeleteBondInfo() error
	StartEncryption() error
	SetWritePDUFunc(func([]byte) (int, error))
	SetEncryptFunc(func(BondInfo) error)
	LegacyPairingInfo() (bool, []byte)
}

type SmpConfig struct {
	IoCap          byte
	OobFlag        byte
	AuthReq        byte
	MaxKeySize     byte
	InitKeyDist    byte
	RespKeyDist    byte
	Authentication ble.AuthData
}

//todo: make these configurable
var defaultSmpConfig = SmpConfig{
	IoCapsKeyboardDisplay, byte(OobNotPresent), 0x09, 16, 0x00, 0x01, ble.AuthData{},
}
