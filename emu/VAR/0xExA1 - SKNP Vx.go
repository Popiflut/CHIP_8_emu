package VAR

// 0xEXA1 SKNP Vx -> Skip next instruction if key with the value of Vx is not pressed.
// Skip next instruction if key with the value of Vx is not pressed. Checks the keyboard, and if the key
// corresponding to the value of Vx is currently in the up position, PC is increased by 2.
func SKNP_Vx(b uint16) {
	if !CHIP8.Clavier.IsPressed[CHIP8.Cpu.V[(b&0x0F00)>>8]] {
		CHIP8.Cpu.Pc += 2
	}
}
