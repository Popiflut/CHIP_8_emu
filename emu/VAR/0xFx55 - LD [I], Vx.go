package VAR

// 0xFX55 LD [I], Vx -> Store registers V0 through Vx in memory starting at location I.
func LD_I_Vx(b uint16) {
	for i := uint16(0); i <= (b&0x0F00)>>8; i++ {
		CHIP8.Cpu.Memory[CHIP8.Cpu.I+i] = CHIP8.Cpu.V[i]
	}
}
