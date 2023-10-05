package VAR

// 0x8XY2 AND Vx, Vy -> Set Vx = Vx AND Vy.
// Set Vx = Vx AND Vy. Performs a bitwise AND on the values of Vx and Vy, then stores the result in Vx.
// A bitwise AND compares the corresponding bits from two values, and if both bits are 1, then the same bit
// in the result is also 1. Otherwise, it is 0.
func AND_Vx_Vy(b uint16) {
	CHIP8.Cpu.V[(b&0x0F00)>>8] &= CHIP8.Cpu.V[(b&0x00F0)>>4]
	CHIP8.Cpu.V[0xF] = 0
}
