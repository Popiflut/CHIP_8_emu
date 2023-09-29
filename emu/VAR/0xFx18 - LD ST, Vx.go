package VAR

// 0xFX18 LD ST, Vx -> Set sound timer = Vx.
// Set sound timer = Vx. Sound Timer is set equal to the value of Vx.
func LD_ST_Vx(b uint16) {
	CHIP8.Cpu.St = CHIP8.Cpu.V[(b&0x0F00)>>8]
}
