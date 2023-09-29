package VAR

// 0x8XY6 SHR Vx {, Vy} -> Set Vx = Vx SHR 1.
func SHR_Vx_Vy(b uint16) {
	CHIP8.Cpu.V[0xF] = CHIP8.Cpu.V[(b&0x0F00)>>8] & 0x1
	CHIP8.Cpu.V[(b&0x0F00)>>8] >>= 1
}
