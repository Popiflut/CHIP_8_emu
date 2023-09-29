package VAR

// 0xFX29 LD F, Vx -> Set I = location of sprite for digit Vx.
func LD_F_Vx(b uint16) {
	CHIP8.Cpu.I = uint16(CHIP8.Cpu.V[(b&0x0F00)>>8] * 5)
}
