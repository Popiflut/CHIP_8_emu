package VAR

// 0x9XY0 SNE Vx, Vy -> Skip next instruction if Vx != Vy.
// Skip next instruction if Vx != Vy. The values of Vx and Vy are compared, and if they are not equal, the
// program counter is increased by 2
func SNE_Vx_Vy(b uint16) {
	if CHIP8.Cpu.V[(b&0x0F00)>>8] != CHIP8.Cpu.V[(b&0x00F0)>>4] {
		CHIP8.Cpu.Pc += 2
	}
}
