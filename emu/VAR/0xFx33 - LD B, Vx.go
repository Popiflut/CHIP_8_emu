package VAR

// 0xFX33 LD B, Vx -> Store BCD representation of Vx in memory locations I, I+1, and I+2.
// Store BCD representation of Vx in memory locations I, I+1, and I+2. The interpreter takes the decimal
// value of Vx, and places the hundreds digit in memory at location in I, the tens digit at location I+1, and
// the ones digit at location I+2.
func LD_B_Vx(b uint16) {
	CHIP8.Cpu.Memory[CHIP8.Cpu.I] = CHIP8.Cpu.V[(b&0x0F00)>>8] / 100
	CHIP8.Cpu.Memory[CHIP8.Cpu.I+1] = (CHIP8.Cpu.V[(b&0x0F00)>>8] / 10) % 10
	CHIP8.Cpu.Memory[CHIP8.Cpu.I+2] = (CHIP8.Cpu.V[(b&0x0F00)>>8] % 100) % 10
}
