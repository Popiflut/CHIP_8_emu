package VAR

// 0x8XYE SHL Vx {, Vy} -> Set Vx = Vx SHL 1.
func SHL_Vx_Vy(b uint16) {
	CHIP8.Cpu.V[0xF] = CHIP8.Cpu.V[(b&0x0F00)>>8] >> 7
	CHIP8.Cpu.V[(b&0x0F00)>>8] <<= 1
}
