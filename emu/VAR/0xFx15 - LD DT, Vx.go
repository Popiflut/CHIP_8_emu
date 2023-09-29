package VAR

// 0xFX15 LD DT, Vx -> Set delay timer = Vx.
func LD_DT_Vx(b uint16) {
	CHIP8.Cpu.Dt = CHIP8.Cpu.V[(b&0x0F00)>>8]
}
