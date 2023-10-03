package VAR

// 0x8XY5 SUB Vx, Vy -> Set Vx = Vx - Vy, set VF = NOT borrow.
// Set Vx = Vx - Vy, set VF = NOT borrow. If Vx ¿ Vy, then VF is set to 1, otherwise 0. Then Vy is
// subtracted from Vx, and the results stored in Vx.
func SUB_Vx_Vy(b uint16) {
	x := (b & 0x0F00) >> 8
	y := (b & 0x00F0) >> 4

	CHIP8.Cpu.V[x] -= CHIP8.Cpu.V[y]
	if CHIP8.Cpu.V[x] > CHIP8.Cpu.V[y] {
		CHIP8.Cpu.V[0xF] = 1
	} else {
		CHIP8.Cpu.V[0xF] = 0
	}
}
