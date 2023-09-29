package VAR

// 0xEX9E SKP Vx -> Skip next instruction if key with the value of Vx is pressed.
func SKP_Vx(b uint16) {
	if CHIP8.Clavier.IsPressed[CHIP8.Cpu.V[(b&0x0F00)>>8]] {
		CHIP8.Cpu.Pc += 2
	}
}
