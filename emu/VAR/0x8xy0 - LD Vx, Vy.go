package VAR

// 0x8XY0 LD Vx, Vy -> Set Vx = Vy.
func LD_Vx_Vy(b uint16) {
	CHIP8.Cpu.V[(b&0x0F00)>>8] = CHIP8.Cpu.V[(b&0x00F0)>>4]
}