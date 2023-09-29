package VAR

// 0x2NNN CALL addr -> Call subroutine at nnn.
func CALL_addr(b uint16) {
	StackPush(CHIP8.Cpu.Pc)
	CHIP8.Cpu.Pc = b&0x0FFF - 2
}
