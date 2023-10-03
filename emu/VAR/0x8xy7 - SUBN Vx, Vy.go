package VAR

import "main/emu/Utils"

// 0x8XY7 SUBN Vx, Vy -> Set Vx = Vy - Vx, set VF = NOT borrow.
// Set Vx = Vy - Vx, set VF = NOT borrow. If Vy ¿ Vx, then VF is set to 1, otherwise 0. Then Vx is
// subtracted from Vy, and the results stored in Vx.
func SUBN_Vx_Vy(b uint16) {
	upper, lower := Utils.Uint16ToUint8(b)
	_, x := Utils.Uint8ToUint4(upper)
	y, _ := Utils.Uint8ToUint4(lower)

	CHIP8.Cpu.V[x] = CHIP8.Cpu.V[y] - CHIP8.Cpu.V[x]
	if CHIP8.Cpu.V[x] < CHIP8.Cpu.V[y] || CHIP8.Cpu.V[x] != CHIP8.Cpu.V[y] {
		CHIP8.Cpu.V[0xF] = 1
	} else {
		CHIP8.Cpu.V[0xF] = 0
	}
}
