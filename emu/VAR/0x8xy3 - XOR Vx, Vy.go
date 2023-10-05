package VAR

// 0x8XY3 XOR Vx, Vy -> Set Vx = Vx XOR Vy.
// Set Vx = Vx XOR Vy. Performs a bitwise exclusive OR on the values of Vx and Vy, then stores the result
// in Vx. An exclusive OR compares the corresponding bits from two values, and if the bits are not both the
// same, then the corresponding bit in the result is set to 1. Otherwise, it is 0.
func XOR_Vx_Vy(b uint16) {
	CHIP8.Cpu.V[(b&0x0F00)>>8] ^= CHIP8.Cpu.V[(b&0x00F0)>>4]
	CHIP8.Cpu.V[0xF] = 0
}
