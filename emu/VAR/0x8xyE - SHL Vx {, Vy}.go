package VAR

// 0x8XYE SHL Vx {, Vy} -> Set Vx = Vx SHL 1.
// Set Vx = Vx SHL 1. If the most-significant bit of Vx is 1, then VF is set to 1, otherwise to 0. Then Vx is
// multiplied by 2.
func SHL_Vx_Vy(b uint16) {
	//x := (b & 0x0F00) >> 8
	//CHIP8.Cpu.V[0xF] = CHIP8.Cpu.V[x] & 0x01
	//CHIP8.Cpu.V[x] >>= 1
	x := (b & 0x0F00) >> 8
	y := (b & 0x00F0) >> 4

	if CHIP8.Cpu.V[x]&0x80 == 0x80 {
		CHIP8.Cpu.V[0xF] = 1
	} else {
		CHIP8.Cpu.V[0xF] = 0
	}
	if (b&0x0F00)>>8 != 0xF {
		CHIP8.Cpu.V[x] = CHIP8.Cpu.V[y] << 1
	}
}
