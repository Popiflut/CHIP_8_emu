package VAR

// 0x1NNN JP addr -> Jump to location nnn.
func JP_addr(b uint16) {
	CHIP8.Cpu.Pc = b&0x0FFF - 2
}
