package VAR

// 0x8XY7 SUBN Vx, Vy -> Set Vx = Vy - Vx, set VF = NOT borrow.
// Set Vx = Vy - Vx, set VF = NOT borrow. If Vy ¿ Vx, then VF is set to 1, otherwise 0. Then Vx is
// subtracted from Vy, and the results stored in Vx.
func SUBN_Vx_Vy(b uint16) {
	x := (b & 0x0F00) >> 8
	y := (b & 0x00F0) >> 4

	CHIP8.Cpu.V[x] = CHIP8.Cpu.V[y] - CHIP8.Cpu.V[x]
	if CHIP8.Cpu.V[x] > CHIP8.Cpu.V[y] {
		CHIP8.Cpu.V[0xF] = 0
	} else {
		CHIP8.Cpu.V[0xF] = 1
	}
}
