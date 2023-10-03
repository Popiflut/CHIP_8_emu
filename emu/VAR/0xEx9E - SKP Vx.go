package VAR

import "main/emu/Utils"

// 0xEX9E SKP Vx -> Skip next instruction if key with the value of Vx is pressed.
// Skip next instruction if key with the value of Vx is pressed. Checks the keyboard, and if the key corresponding
// to the value of Vx is currently in the down position, PC is increased by 2.
func SKP_Vx(b uint16) {
	a, _ := Utils.Uint16ToUint8(b)
	_, x := Utils.Uint8ToUint4(a)

	if CHIP8.Clavier.IsPressed[CHIP8.Cpu.V[x]] {
		CHIP8.Cpu.Pc += 2
	}
}
