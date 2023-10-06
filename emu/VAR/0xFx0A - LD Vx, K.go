package VAR

// 0xFX0A LD Vx, K -> Wait for a key press, store the value of the key in Vx.
// Wait for a key press, store the value of the key in Vx. All execution stops until a key is pressed, then the
// value of that key is stored in Vx.
func LD_Vx_K(b uint16) {
	var key bool
	x := (b & 0x0F00) >> 8
	for i, _ := range CHIP8.Clavier.IsPressed {
		if CHIP8.Clavier.IsPressed[i] {
			key = true
			CHIP8.Cpu.V[x] = uint8(i)
		}
	}
	if !key {
		CHIP8.Cpu.Pc -= 2
	}
}
