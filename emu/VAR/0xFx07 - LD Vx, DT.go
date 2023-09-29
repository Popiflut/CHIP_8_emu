package VAR

// 0xFX07 LD Vx, DT -> Set Vx = delay timer value.
// Set Vx = delay timer value. The value of DT is placed into Vx.
func LD_Vx_DT(b uint16) {
	CHIP8.Cpu.V[(b&0x0F00)>>8] = CHIP8.Cpu.Dt
}
