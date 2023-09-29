package OpCodes

import "main/emu"

// 0x000E RET -> Return from a subroutine.
func RET() {
	emu.CHIP8.Cpu.Pc = emu.StackPop()
}
