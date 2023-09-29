package VAR

// 0xFX29 LD F, Vx -> Set I = location of sprite for digit Vx.
// Set I = location of sprite for digit Vx. The value of I is set to the location for the hexadecimal sprite
// corresponding to the value of Vx. See section 2.4, Display, for more information on the Chip-8 hexadecimal
// font. To obtain this value, multiply VX by 5 (all font data stored in first 80 bytes of memory).
func LD_F_Vx(b uint16) {
	CHIP8.Cpu.I = uint16(CHIP8.Cpu.V[(b&0x0F00)>>8] * 5)
}
