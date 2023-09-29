package VAR

// 0x000E RET -> Return from a subroutine.
// Return from a subroutine.The interpreter sets the program counter to the address at the top of the stack,
// then subtracts 1 from the stack pointer.
func RET() {
	CHIP8.Cpu.Pc = StackPop()
}
