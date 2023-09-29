package VAR

// 0x3XNN SE Vx, byte -> Skip next instruction if Vx = kk.
func SE_Vx_byte(b uint16) {
	if CHIP8.Cpu.V[(b&0x0F00)>>8] == uint8(b&0x00FF) {
		CHIP8.Cpu.Pc += 2
	}
}