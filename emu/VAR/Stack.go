package VAR

func StackPush(pc uint16) {
	CHIP8.Cpu.Stack[CHIP8.Cpu.Sp] = pc
	CHIP8.Cpu.Sp++
}

func StackPop() uint16 {
	CHIP8.Cpu.Sp--
	return CHIP8.Cpu.Stack[CHIP8.Cpu.Sp]
}
