package VAR

// 0x8XY4 ADD Vx, Vy -> Set Vx = Vx + Vy, set VF = carry.
func ADD_Vx_Vy(b uint16) {
	if CHIP8.Cpu.V[(b&0x00F0)>>4] > (0xFF - CHIP8.Cpu.V[(b&0x0F00)>>8]) {
		CHIP8.Cpu.V[0xF] = 1
	} else {
		CHIP8.Cpu.V[0xF] = 0
	}
	CHIP8.Cpu.V[(b&0x0F00)>>8] += CHIP8.Cpu.V[(b&0x00F0)>>4]
}
