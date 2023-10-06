package VAR

import (
	"main/emu/Utils"
	"math/rand"
)

// 0x00E0 CLS -> Clear the display.
// Clear the display.
func CLS() {
	for i := 0; i < ResolWidth; i++ {
		for j := 0; j < ResolHeight; j++ {
			CHIP8.Screen.Mapscreen[i][j] = 0
		}
	}
}

// 0x00EE RET -> Return from a subroutine.
// Return from a subroutine.The interpreter sets the program counter to the address at the top of the stack,
// then subtracts 1 from the stack pointer.
func RET() {
	CHIP8.Cpu.Pc = StackPop()
}

// 0x1NNN JP addr -> Jump to location nnn.
// Jump to location nnn. The interpreter sets the program counter to nnn.
func JP_addr(b uint16) {
	CHIP8.Cpu.Pc = b&0x0FFF - 2
}

// 0x2NNN CALL addr -> Call subroutine at nnn.
// The interpreter increments the stack pointer, then puts the current PC on the top of the stack. The PC is then set to nnn.
func CALL_addr(b uint16) {
	StackPush(CHIP8.Cpu.Pc)
	CHIP8.Cpu.Pc = b&0x0FFF - 2
}

// 0x3XNN SE Vx, byte -> Skip next instruction if Vx = kk.
// Skip next instruction if Vx = kk. The interpreter compares register Vx to kk, and if they are equal,
// increments the program counter by 2.
func SE_Vx_byte(b uint16) {
	if CHIP8.Cpu.V[(b&0x0F00)>>8] == uint8(b&0x00FF) {
		CHIP8.Cpu.Pc += 2
	}
}

// 0x4XNN SNE Vx, byte -> Skip next instruction if Vx != kk.
// Skip next instruction if Vx != kk. The interpreter compares register Vx to kk, and if they are not equal,
// increments the program counter by 2.
func SNE_Vx_byte(b uint16) {
	if CHIP8.Cpu.V[(b&0x0F00)>>8] != uint8(b&0x00FF) {
		CHIP8.Cpu.Pc += 2
	}
}

// 0x5XY0 SE Vx, Vy -> Skip next instruction if Vx = Vy.
// Skip next instruction if Vx = Vy. The interpreter compares register Vx to register Vy, and if they are equal,
// increments the program counter by 2.
func SE_Vx_Vy(b uint16) {
	if CHIP8.Cpu.V[(b&0x0F00)>>8] == CHIP8.Cpu.V[(b&0x00F0)>>4] {
		CHIP8.Cpu.Pc += 2
	}
}

// 0x6XNN LD Vx, byte -> Set Vx = kk.
// Set Vx = kk. The interpreter puts the value kk into register Vx.
func LD_Vx_byte(b uint16) {
	CHIP8.Cpu.V[(b&0x0F00)>>8] = uint8(b & 0x00FF)
}

// 0x7XNN ADD Vx, byte -> Set Vx = Vx + kk.
// Set Vx = Vx + kk. Adds the value kk to the value of register Vx, then stores the result in Vx
func ADD_Vx_byte(b uint16) {
	CHIP8.Cpu.V[(b&0x0F00)>>8] += uint8(b & 0x00FF)
}

// 0x8XY0 LD Vx, Vy -> Set Vx = Vy.
// Set Vx = Vy. Stores the value of register Vy in register Vx.
func LD_Vx_Vy(b uint16) {
	CHIP8.Cpu.V[(b&0x0F00)>>8] = CHIP8.Cpu.V[(b&0x00F0)>>4]
}

// 0x8XY1 OR Vx, Vy -> Set Vx = Vx OR Vy.
// Set Vx = Vx OR Vy. Performs a bitwise OR on the values of Vx and Vy, then stores the result in Vx. A
// bitwise OR compares the corresponding bits from two values, and if either bit is 1, then the same bit in the
// result is also 1. Otherwise, it is 0.
func OR_Vx_Vy(b uint16) {
	CHIP8.Cpu.V[(b&0x0F00)>>8] |= CHIP8.Cpu.V[(b&0x00F0)>>4]
	CHIP8.Cpu.V[0xF] = 0
}

// 0x8XY2 AND Vx, Vy -> Set Vx = Vx AND Vy.
// Set Vx = Vx AND Vy. Performs a bitwise AND on the values of Vx and Vy, then stores the result in Vx.
// A bitwise AND compares the corresponding bits from two values, and if both bits are 1, then the same bit
// in the result is also 1. Otherwise, it is 0.
func AND_Vx_Vy(b uint16) {
	CHIP8.Cpu.V[(b&0x0F00)>>8] &= CHIP8.Cpu.V[(b&0x00F0)>>4]
	CHIP8.Cpu.V[0xF] = 0
}

// 0x8XY3 XOR Vx, Vy -> Set Vx = Vx XOR Vy.
// Set Vx = Vx XOR Vy. Performs a bitwise exclusive OR on the values of Vx and Vy, then stores the result
// in Vx. An exclusive OR compares the corresponding bits from two values, and if the bits are not both the
// same, then the corresponding bit in the result is set to 1. Otherwise, it is 0.
func XOR_Vx_Vy(b uint16) {
	CHIP8.Cpu.V[(b&0x0F00)>>8] ^= CHIP8.Cpu.V[(b&0x00F0)>>4]
	CHIP8.Cpu.V[0xF] = 0
}

// 0x8XY4 ADD Vx, Vy -> Set Vx = Vx + Vy, set VF = carry.
// Set Vx = Vx + Vy, set VF = carry. The values of Vx and Vy are added together. If the result is greater
// than 8 bits (i.e., ¿ 255,) VF is set to 1, otherwise 0. Only the lowest 8 bits of the result are kept, and stored
// in Vx.
func ADD_Vx_Vy(b uint16) {
	x := (b & 0x0F00) >> 8
	y := (b & 0x00F0) >> 4

	a := uint16(CHIP8.Cpu.V[x]) + uint16(CHIP8.Cpu.V[y])
	if a > 255 {
		CHIP8.Cpu.V[0xF] = 0x1
	} else {
		CHIP8.Cpu.V[0xF] = 0x0
	}
	if x != 0xF {
		CHIP8.Cpu.V[x] = byte(a)
	}
}

// 0x8XY5 SUB Vx, Vy -> Set Vx = Vx - Vy, set VF = NOT borrow.
// Set Vx = Vx - Vy, set VF = NOT borrow. If Vx ¿ Vy, then VF is set to 1, otherwise 0. Then Vy is
// subtracted from Vx, and the results stored in Vx.
func SUB_Vx_Vy(b uint16) {
	x := (b & 0x0F00) >> 8
	y := (b & 0x00F0) >> 4

	a := int16(CHIP8.Cpu.V[x]) - int16(CHIP8.Cpu.V[y])
	if a < 0 {
		CHIP8.Cpu.V[0xF] = 0x0
	} else {
		CHIP8.Cpu.V[0xF] = 0x1
	}
	if x != 0xF {
		CHIP8.Cpu.V[x] = byte(a)
	}
}

// 0x8XY6 SHR Vx {, Vy} -> Set Vx = Vx SHR 1.
// Set Vx = Vx SHR 1. If the least-significant bit of Vx is 1, then VF is set to 1, otherwise 0. Then Vx is
// divided by 2.
func SHR_Vx_Vy(b uint16) {
	//x := (b & 0x0F00) >> 8
	//CHIP8.Cpu.V[0xF] = CHIP8.Cpu.V[x] & 0x01
	//CHIP8.Cpu.V[x] >>= 1
	x := (b & 0x0F00) >> 8
	y := (b & 0x00F0) >> 4

	if CHIP8.Cpu.V[x]&0x01 == 0x01 {
		CHIP8.Cpu.V[0xF] = 1
	} else {
		CHIP8.Cpu.V[0xF] = 0
	}
	if (b&0x0F00)>>8 != 0xF {
		CHIP8.Cpu.V[x] = CHIP8.Cpu.V[y] >> 1
	}
}

// 0x8XY7 SUBN Vx, Vy -> Set Vx = Vy - Vx, set VF = NOT borrow.
// Set Vx = Vy - Vx, set VF = NOT borrow. If Vy ¿ Vx, then VF is set to 1, otherwise 0. Then Vx is
// subtracted from Vy, and the results stored in Vx.
func SUBN_Vx_Vy(b uint16) {
	x := (b & 0x0F00) >> 8
	y := (b & 0x00F0) >> 4

	CHIP8.Cpu.V[x] = CHIP8.Cpu.V[y] - CHIP8.Cpu.V[x]
	if CHIP8.Cpu.V[x] > CHIP8.Cpu.V[y] {
		CHIP8.Cpu.V[0xF] = 0
	} else {
		CHIP8.Cpu.V[0xF] = 1
	}
}

// 0x8XYE SHL Vx {, Vy} -> Set Vx = Vx SHL 1.
// Set Vx = Vx SHL 1. If the most-significant bit of Vx is 1, then VF is set to 1, otherwise to 0. Then Vx is
// multiplied by 2.
func SHL_Vx_Vy(b uint16) {
	//x := (b & 0x0F00) >> 8
	//CHIP8.Cpu.V[0xF] = CHIP8.Cpu.V[x] & 0x01
	//CHIP8.Cpu.V[x] >>= 1
	x := (b & 0x0F00) >> 8
	y := (b & 0x00F0) >> 4

	if CHIP8.Cpu.V[x]&0x80 == 0x80 {
		CHIP8.Cpu.V[0xF] = 1
	} else {
		CHIP8.Cpu.V[0xF] = 0
	}
	if (b&0x0F00)>>8 != 0xF {
		CHIP8.Cpu.V[x] = CHIP8.Cpu.V[y] << 1
	}
}

// 0x9XY0 SNE Vx, Vy -> Skip next instruction if Vx != Vy.
// Skip next instruction if Vx != Vy. The values of Vx and Vy are compared, and if they are not equal, the
// program counter is increased by 2
func SNE_Vx_Vy(b uint16) {
	if CHIP8.Cpu.V[(b&0x0F00)>>8] != CHIP8.Cpu.V[(b&0x00F0)>>4] {
		CHIP8.Cpu.Pc += 2
	}
}

// 0xANNN LD I, addr -> Set I = nnn.
// Set I = nnn. The value of register I is set to nnn.
func LD_I_addr(b uint16) {
	CHIP8.Cpu.I = b & 0x0FFF
}

// 0xBNNN JP V0, addr -> Jump to location nnn + V0.
// Jump to location nnn + V0. The program counter is set to nnn plus the value of V0.
func JP_V0_addr(b uint16) {
	CHIP8.Cpu.Pc = (b & 0x0FFF) + uint16(CHIP8.Cpu.V[b&0x000F>>8])
}

// 0xCXNN RND Vx, byte -> Set Vx = random byte AND nn.
// Set Vx = random byte AND kk. The interpreter generates a random number from 0 to 255, which is then
// ANDed with the value kk. The results are stored in Vx. See instruction 8xy2 for more information on AND.
func RND_Vx_byte(b uint16) {
	CHIP8.Cpu.V[(b&0x0F00)>>8] = uint8(rand.Intn(255)) & uint8(b&0x00FF)
}

// 0xDXYK DRW Vx, Vy, nibble -> Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision.
// Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision. The interpreter reads n
// bytes from memory, starting at the address stored in I. These bytes are then displayed as sprites on screen
// at coordinates (Vx, Vy). Sprites are XOR’d onto the existing screen. If this causes any pixels to be erased,
// VF is set to 1, otherwise it is set to 0. If the sprite is positioned so part of it is outside the coordinates of
// the display, it wraps around to the opposite side of the screen.
func DRW_Vx_Vy_nibble(b uint16) {
	//var tmps []uint8
	//for i := CHIP8.Cpu.I; i < CHIP8.Cpu.I+b&0x000F; i++ {
	//	tmps = append(tmps, CHIP8.Cpu.Memory[i])
	//}
	//for i := uint8(0); i < uint8(len(tmps)); i++ {
	//	for j := uint8(0); j < 8; j++ {
	//		CHIP8.Screen.Mapscreen[(CHIP8.Cpu.V[(b&0x0F00)>>8]+j)%64][(CHIP8.Cpu.V[(b&0x00F0)>>4]+i)%32] ^= tmps[i] >> (7 - j) & 0x01
	//	}
	//}

	xval := CHIP8.Cpu.V[b&0x0F00>>8]
	yval := CHIP8.Cpu.V[b&0x00F0>>4]
	CHIP8.Cpu.V[0xF] = 0
	var i byte = 0
	for ; i < byte(b&0x000F); i++ {
		row := CHIP8.Cpu.Memory[CHIP8.Cpu.I+uint16(i)]
		if erased := DrawSprite(xval, yval+i, row); erased {
			CHIP8.Cpu.V[0xF] = 1
		}
	}

}

func DrawSprite(x byte, y byte, row byte) bool {
	erased := false
	yIndex := y % 32
	for i := x; i < x+8; i++ {
		xIndex := i % 64

		wasSet := CHIP8.Screen.Mapscreen[xIndex][yIndex] == 1
		value := row >> (x + 8 - i - 1) & 0x01

		CHIP8.Screen.Mapscreen[xIndex][yIndex] ^= value

		if wasSet && CHIP8.Screen.Mapscreen[xIndex][yIndex] == 0 {
			erased = true
		}
	}
	return erased
}

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

// 0xEXA1 SKNP Vx -> Skip next instruction if key with the value of Vx is not pressed.
// Skip next instruction if key with the value of Vx is not pressed. Checks the keyboard, and if the key
// corresponding to the value of Vx is currently in the up position, PC is increased by 2.
func SKNP_Vx(b uint16) {
	a, _ := Utils.Uint16ToUint8(b)
	_, x := Utils.Uint8ToUint4(a)

	if !CHIP8.Clavier.IsPressed[CHIP8.Cpu.V[x]] {
		CHIP8.Cpu.Pc += 2
	}
}

// 0xFX0A LD Vx, K -> Wait for a key press, store the value of the key in Vx.
// Wait for a key press, store the value of the key in Vx. All execution stops until a key is pressed, then the
// value of that key is stored in Vx.
func LD_Vx_K(b uint16) {
	var key bool
	x := (b & 0x0F00) >> 8
	for i := range CHIP8.Clavier.IsPressed {
		if CHIP8.Clavier.IsPressed[i] {
			key = true
			CHIP8.Cpu.V[x] = uint8(i)
		}
	}
	if !key {
		CHIP8.Cpu.Pc -= 2
	}
}

// 0xFX1E ADD I, Vx -> Set I = I + Vx.
// Set I = I + Vx. The values of I and Vx are added, and the results are stored in I.
func ADD_I_Vx(b uint16) {
	CHIP8.Cpu.I += uint16(CHIP8.Cpu.V[(b&0x0F00)>>8])
}

// 0xFX07 LD Vx, DT -> Set Vx = delay timer value.
// Set Vx = delay timer value. The value of DT is placed into Vx.
func LD_Vx_DT(b uint16) {
	CHIP8.Cpu.V[(b&0x0F00)>>8] = CHIP8.Cpu.Dt
}

// 0xFX15 LD DT, Vx -> Set delay timer = Vx.
// Set delay timer = Vx. Delay Timer is set equal to the value of Vx.
func LD_DT_Vx(b uint16) {
	CHIP8.Cpu.Dt = CHIP8.Cpu.V[(b&0x0F00)>>8]
}

// 0xFX18 LD ST, Vx -> Set sound timer = Vx.
// Set sound timer = Vx. Sound Timer is set equal to the value of Vx.
func LD_ST_Vx(b uint16) {
	CHIP8.Cpu.SoundTimer = CHIP8.Cpu.V[(b&0x0F00)>>8]
}

// 0xFX29 LD F, Vx -> Set I = location of sprite for digit Vx.
// Set I = location of sprite for digit Vx. The value of I is set to the location for the hexadecimal sprite
// corresponding to the value of Vx. See section 2.4, Display, for more information on the Chip-8 hexadecimal
// font. To obtain this value, multiply VX by 5 (all font data stored in first 80 bytes of memory).
func LD_F_Vx(b uint16) {
	CHIP8.Cpu.I = uint16(CHIP8.Cpu.V[(b&0x0F00)>>8] * 5)
}

// 0xFX33 LD B, Vx -> Store BCD representation of Vx in memory locations I, I+1, and I+2.
// Store BCD representation of Vx in memory locations I, I+1, and I+2. The interpreter takes the decimal
// value of Vx, and places the hundreds digit in memory at location in I, the tens digit at location I+1, and
// the ones digit at location I+2.
func LD_B_Vx(b uint16) {
	CHIP8.Cpu.Memory[CHIP8.Cpu.I] = CHIP8.Cpu.V[(b&0x0F00)>>8] / 100
	CHIP8.Cpu.Memory[CHIP8.Cpu.I+1] = (CHIP8.Cpu.V[(b&0x0F00)>>8] / 10) % 10
	CHIP8.Cpu.Memory[CHIP8.Cpu.I+2] = (CHIP8.Cpu.V[(b&0x0F00)>>8] % 100) % 10
}

// 0xFX55 LD [I], Vx -> Store registers V0 through Vx in memory starting at location I.
// Stores V0 to VX in memory starting at address I. I is then set to I + x + 1.
func LD_I_Vx(b uint16) {
	for i := uint16(0); i <= (b&0x0F00)>>8; i++ {
		CHIP8.Cpu.Memory[CHIP8.Cpu.I+i] = CHIP8.Cpu.V[i]
	}
	CHIP8.Cpu.I += b&0x0F00>>8 + 1

	//for i := uint16(0); i <= (b&0x0F00)>>8; i++ {
	//	CHIP8.Cpu.Memory[CHIP8.Cpu.I+i] = CHIP8.Cpu.V[i]
	//}
}

// 0xFX65 LD Vx, [I] -> Read registers V0 through Vx from memory starting at location I.
// Fills V0 to VX with values from memory starting at address I. I is then set to I + x + 1.
func LD_Vx_I(b uint16) {
	for i := uint16(0); i <= b&0x0F00>>8; i++ {
		CHIP8.Cpu.V[i] = CHIP8.Cpu.Memory[CHIP8.Cpu.I+i]
	}
	CHIP8.Cpu.I += b&0x0F00>>8 + 1
}
