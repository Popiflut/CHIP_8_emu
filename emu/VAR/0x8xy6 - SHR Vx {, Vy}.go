package VAR

// 0x8XY6 SHR Vx {, Vy} -> Set Vx = Vx SHR 1.
// Set Vx = Vx SHR 1. If the least-significant bit of Vx is 1, then VF is set to 1, otherwise 0. Then Vx is
// divided by 2.
func SHR_Vx_Vy(b uint16) {
	//x := (b & 0x0F00) >> 8
	//CHIP8.Cpu.V[0xF] = CHIP8.Cpu.V[x] & 0x01
	//CHIP8.Cpu.V[x] >>= 1
	x := (b & 0x0F00) >> 8
	y := (b & 0x00F0) >> 4

	if CHIP8.Cpu.V[x]&0x01 == 0x01 {
		CHIP8.Cpu.V[0xF] = 1
	} else {
		CHIP8.Cpu.V[0xF] = 0
	}
	if (b&0x0F00)>>8 != 0xF {
		CHIP8.Cpu.V[x] = CHIP8.Cpu.V[y] >> 1
	}
}
