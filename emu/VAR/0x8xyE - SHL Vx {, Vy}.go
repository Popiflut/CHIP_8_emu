package VAR

// 0x8XYE SHL Vx {, Vy} -> Set Vx = Vx SHL 1.
// Set Vx = Vx SHL 1. If the most-significant bit of Vx is 1, then VF is set to 1, otherwise to 0. Then Vx is
// multiplied by 2.
func SHL_Vx_Vy(b uint16) {
	CHIP8.Cpu.V[0xF] = CHIP8.Cpu.V[(b&0x0F00)>>8] >> 7
	CHIP8.Cpu.V[(b&0x0F00)>>8] <<= 1
}
