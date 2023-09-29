package VAR

// 0xFX0A LD Vx, K -> Wait for a key press, store the value of the key in Vx.
func LD_Vx_K(b uint16) {
	CHIP8.Cpu.V[(b&0x0F00)>>8] = CHIP8.Cpu.Dt
}
