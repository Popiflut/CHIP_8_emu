package VAR

// 0x8XY5 SUB Vx, Vy -> Set Vx = Vx - Vy, set VF = NOT borrow.
func SUB_Vx_Vy(b uint16) {
	if CHIP8.Cpu.V[(b&0x00F0)>>4] > CHIP8.Cpu.V[(b&0x0F00)>>8] {
		CHIP8.Cpu.V[0xF] = 0
	} else {
		CHIP8.Cpu.V[0xF] = 1
	}
	CHIP8.Cpu.V[(b&0x0F00)>>8] -= CHIP8.Cpu.V[(b&0x00F0)>>4]
}
