package VAR

import "main/emu/Utils"

// 0x8XY4 ADD Vx, Vy -> Set Vx = Vx + Vy, set VF = carry.
// Set Vx = Vx + Vy, set VF = carry. The values of Vx and Vy are added together. If the result is greater
// than 8 bits (i.e., ¿ 255,) VF is set to 1, otherwise 0. Only the lowest 8 bits of the result are kept, and stored
// in Vx.
func ADD_Vx_Vy(b uint16) {
	upper, lower := Utils.Uint16ToUint8(b)
	_, x := Utils.Uint8ToUint4(upper)
	y, _ := Utils.Uint8ToUint4(lower)

	CHIP8.Cpu.V[x] += CHIP8.Cpu.V[y]
	if CHIP8.Cpu.V[x] >= 0xFF {
		CHIP8.Cpu.V[0xF] = 1
	} else {
		CHIP8.Cpu.V[0xF] = 0
	}
}
