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
	AVR  Core = 0
	AVRe Core = 1
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
		AVR,
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
