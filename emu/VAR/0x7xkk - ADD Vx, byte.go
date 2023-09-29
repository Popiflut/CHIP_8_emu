package VAR

// 0x7XNN ADD Vx, byte -> Set Vx = Vx + kk.
func ADD_Vx_byte(b uint16) {
	CHIP8.Cpu.V[(b&0x0F00)>>8] += uint8(b & 0x00FF)
}
