package VAR

import "fmt"

// 0x4XNN SNE Vx, byte -> Skip next instruction if Vx != kk.
// Skip next instruction if Vx != kk. The interpreter compares register Vx to kk, and if they are not equal,
// increments the program counter by 2.
func SNE_Vx_byte(b uint16) {
	if CHIP8.Cpu.V[(b&0x0F00)>>8] != uint8(b&0x00FF) {
		fmt.Printf("\n SNE V%X != %02X\n", (b&0x0F00)>>8, b&0x00FF)
		CHIP8.Cpu.Pc += 2
	}
}
