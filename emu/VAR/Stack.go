package VAR

// StackPush pushes the current PC to the stack
func StackPush(pc uint16) {
	CHIP8.Cpu.Stack[CHIP8.Cpu.Sp] = pc
	CHIP8.Cpu.Sp++
}

// StackPop return to the current PC the last pushed from the stack
func StackPop() uint16 {
	CHIP8.Cpu.Sp--
	return CHIP8.Cpu.Stack[CHIP8.Cpu.Sp]
}
