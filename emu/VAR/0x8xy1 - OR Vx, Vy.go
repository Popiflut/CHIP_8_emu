package VAR

// 0x8XY1 OR Vx, Vy -> Set Vx = Vx OR Vy.
// Set Vx = Vx OR Vy. Performs a bitwise OR on the values of Vx and Vy, then stores the result in Vx. A
// bitwise OR compares the corresponding bits from two values, and if either bit is 1, then the same bit in the
// result is also 1. Otherwise, it is 0.
func OR_Vx_Vy(b uint16) {
	CHIP8.Cpu.V[(b&0x0F00)>>8] |= CHIP8.Cpu.V[(b&0x00F0)>>4]
}
