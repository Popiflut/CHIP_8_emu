package Utils

func Uint8ToUint16(a uint8, b uint8) uint16 {
	return uint16(a)<<8 | uint16(b) //8 bits vers la gauche
}

func Uint16ToUint8(a uint16) (uint8, uint8) {
	return uint8(a >> 8), uint8(a & 0x00FF) //8 bits vers la droite
}

func Uint8ToUint4(a uint8) (uint8, uint8) {
	return a >> 4, a & 0x0F //4 bits vers la droite
}

func Uint4ToUint8(a uint8, b uint8) uint8 {
	return uint8(a)<<4 | uint8(b) //4 bits vers la gauche
}

func AtoI(a string) int {
	var tmp int
	for i := 0; i < len(a); i++ {
		tmp += int(a[i]) - 48
		if i != len(a)-1 {
			tmp *= 10
		}
	}
	return tmp
}
