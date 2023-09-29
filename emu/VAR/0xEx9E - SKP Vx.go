package VAR

// 0xEX9E SKP Vx -> Skip next instruction if key with the value of Vx is pressed.
// Skip next instruction if key with the value of Vx is pressed. Checks the keyboard, and if the key corresponding
// to the value of Vx is currently in the down position, PC is increased by 2.
func SKP_Vx(b uint16) {
	if CHIP8.Clavier.IsPressed[CHIP8.Cpu.V[(b&0x0F00)>>8]] {
		CHIP8.Cpu.Pc += 2
	}
}
