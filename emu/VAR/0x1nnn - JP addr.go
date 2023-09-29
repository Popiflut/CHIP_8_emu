package VAR

// 0x1NNN JP addr -> Jump to location nnn.
// Jump to location nnn. The interpreter sets the program counter to nnn.
func JP_addr(b uint16) {
	CHIP8.Cpu.Pc = b&0x0FFF - 2
}
