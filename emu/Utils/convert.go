package Utils

func Uint8ToUint16(a uint8, b uint8) uint16 {
	return uint16(a)<<8 | uint16(b)
}

func Uint16ToUint8(a uint16) (uint8, uint8) {
	return uint8(a >> 8), uint8(a & 0x00FF)
}

func Uint8ToUint4(a uint8) (uint8, uint8) {
	return a >> 4, a & 0x0F
}

func Uint4ToUint8(a uint8, b uint8) uint8 {
	return uint8(a)<<4 | uint8(b)
}
