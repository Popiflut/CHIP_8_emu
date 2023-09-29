package VAR

// 0xBNNN JP V0, addr -> Jump to location nnn + V0.
func JP_V0_addr(b uint16) {
	CHIP8.Cpu.Pc = (b & 0x0FFF) + uint16(CHIP8.Cpu.V[0])
}
