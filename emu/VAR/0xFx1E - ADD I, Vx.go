package VAR

// 0xFX1E ADD I, Vx -> Set I = I + Vx.
func ADD_I_Vx(b uint16) {
	CHIP8.Cpu.I += uint16(CHIP8.Cpu.V[(b&0x0F00)>>8])
}
