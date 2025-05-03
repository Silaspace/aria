package device

type DeviceType string
type Flag int
type Core int

type Device struct {
	DeviceCore Core
	RAMStart   uint32
	RAMSize    uint32 /* Bytes */
	EEPROMSize uint32 /* Bytes */
	FlashSize  uint32 /* Words */
	Flags      Flag
}

/* Device cores */
const (
	Nil  Core = 0
	AVR  Core = 1
	AVRe Core = 2
)

/* Device Names */
const (
	DEFAULT    DeviceType = "" /* Based on the attiny25 */
	AT90USB82  DeviceType = "at90usb82"
	AT90USB162 DeviceType = "at90usb162"
)

/* Devices */
var DeviceMap = map[DeviceType]Device{
	DEFAULT: {
		Nil,
		0x060,
		128,
		128,
		2048,
		0,
	},
	AT90USB82: {
		AVRe,
		0x100,
		512,
		512,
		4096,
		0,
	},
	AT90USB162: {
		AVRe,
		0x100,
		512,
		512,
		8192,
		0,
	},
}
