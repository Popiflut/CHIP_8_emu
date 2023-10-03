package VAR

import "fmt"

// 0x3XNN SE Vx, byte -> Skip next instruction if Vx = kk.
// Skip next instruction if Vx = kk. The interpreter compares register Vx to kk, and if they are equal,
// increments the program counter by 2.
func SE_Vx_byte(b uint16) {
	if CHIP8.Cpu.V[(b&0x0F00)>>8] == uint8(b&0x00FF) {
		fmt.Printf("\n SE V%X == %02X\n", (b&0x0F00)>>8, b&0x00FF)
		CHIP8.Cpu.Pc += 2
	}
}
