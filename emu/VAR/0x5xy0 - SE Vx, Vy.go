package VAR

// 0x5XY0 SE Vx, Vy -> Skip next instruction if Vx = Vy.
func SE_Vx_Vy(b uint16) {
	if CHIP8.Cpu.V[(b&0x0F00)>>8] == CHIP8.Cpu.V[(b&0x00F0)>>4] {
		CHIP8.Cpu.Pc += 2
	}
}
