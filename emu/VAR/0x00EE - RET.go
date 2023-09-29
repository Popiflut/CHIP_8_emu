package VAR

// 0x000E RET -> Return from a subroutine.
func RET() {
	CHIP8.Cpu.Pc = StackPop()
}
