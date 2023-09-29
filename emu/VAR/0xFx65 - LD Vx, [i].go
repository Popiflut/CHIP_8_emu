package VAR

// 0xFX65 LD Vx, [I] -> Read registers V0 through Vx from memory starting at location I.
func LD_Vx_I(b uint16) {
	for i := uint16(0); i <= (b&0x0F00)>>8; i++ {
		CHIP8.Cpu.V[i] = CHIP8.Cpu.Memory[CHIP8.Cpu.I+i]
	}
}