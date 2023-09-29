package VAR

// 0x6XNN LD Vx, byte -> Set Vx = kk.
func LD_Vx_byte(b uint16) {
	CHIP8.Cpu.V[(b&0x0F00)>>8] = uint8(b & 0x00FF)
}
