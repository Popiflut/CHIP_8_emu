package VAR

// 0x8XY2 AND Vx, Vy -> Set Vx = Vx AND Vy.
func AND_Vx_Vy(b uint16) {
	CHIP8.Cpu.V[(b&0x0F00)>>8] &= CHIP8.Cpu.V[(b&0x00F0)>>4]
}
