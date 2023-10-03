package VAR

import "main/emu/Utils"

// 0x8XY5 SUB Vx, Vy -> Set Vx = Vx - Vy, set VF = NOT borrow.
// Set Vx = Vx - Vy, set VF = NOT borrow. If Vx ¿ Vy, then VF is set to 1, otherwise 0. Then Vy is
// subtracted from Vx, and the results stored in Vx.
func SUB_Vx_Vy(b uint16) {
	upper, lower := Utils.Uint16ToUint8(b)
	_, x := Utils.Uint8ToUint4(upper)
	y, _ := Utils.Uint8ToUint4(lower)

	CHIP8.Cpu.V[x] -= CHIP8.Cpu.V[y]
	if CHIP8.Cpu.V[x] > CHIP8.Cpu.V[y] || CHIP8.Cpu.V[x] != CHIP8.Cpu.V[y] {
		CHIP8.Cpu.V[0xF] = 1
	} else {
		CHIP8.Cpu.V[0xF] = 0
	}
}
