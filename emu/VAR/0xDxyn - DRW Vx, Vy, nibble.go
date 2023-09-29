package VAR

// 0xDXYK DRW Vx, Vy, nibble -> Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision.
// Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision. The interpreter reads n
// bytes from memory, starting at the address stored in I. These bytes are then displayed as sprites on screen
// at coordinates (Vx, Vy). Sprites are XOR’d onto the existing screen. If this causes any pixels to be erased,
// VF is set to 1, otherwise it is set to 0. If the sprite is positioned so part of it is outside the coordinates of
// the display, it wraps around to the opposite side of the screen.
func DRW_Vx_Vy_nibble(b uint16) {
	var tmps []uint8
	for i := CHIP8.Cpu.I; i < CHIP8.Cpu.I+b&0x000F; i++ {
		tmps = append(tmps, CHIP8.Cpu.Memory[i])
	}
	for i := uint8(0); i < uint8(len(tmps)); i++ {
		for j := uint8(0); j < 8; j++ {
			CHIP8.Screen.Mapscreen[(CHIP8.Cpu.V[(b&0x0F00)>>8]+j)%64][(CHIP8.Cpu.V[(b&0x00F0)>>4]+i)%32] ^= tmps[i] >> (7 - j) & 0x01
		}
	}
}
