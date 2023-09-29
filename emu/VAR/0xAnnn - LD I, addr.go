package VAR

// 0xANNN LD I, addr -> Set I = nnn.
func LD_I_addr(b uint16) {
	CHIP8.Cpu.I = b & 0x0FFF
}
