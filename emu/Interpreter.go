package emu

import (
	"fmt"
	"main/emu/OpCodes"
	"math/rand"
)

func (cpu *CPU) Interpreter(b uint16) {
	switch b & 0xF000 {
	case 0x0000:
		switch b & 0x000F {
		case 0x0000:
			OpCodes.CLS()
		case 0x000E:
			OpCodes.RET()
		}
	case 0x1000:
		//0x1NNN JP addr -> Jump to location nnn.
		CHIP8.Cpu.Pc = b&0x0FFF - 2
	case 0x2000:
		//0x2NNN CALL addr -> Call subroutine at nnn.
		StackPush(CHIP8.Cpu.Pc)
		CHIP8.Cpu.Pc = b&0x0FFF - 2
	case 0x3000:
		//0x3XNN SE Vx, byte -> Skip next instruction if Vx = kk.
		if CHIP8.Cpu.v[(b&0x0F00)>>8] == uint8(b&0x00FF) {
			CHIP8.Cpu.Pc += 2
		}
	case 0x4000:
		//0x4XNN SNE Vx, byte -> Skip next instruction if Vx != kk.
		if CHIP8.Cpu.v[(b&0x0F00)>>8] != uint8(b&0x00FF) {
			CHIP8.Cpu.Pc += 2
		}
	case 0x5000:
		//0x5XY0 SE Vx, Vy -> Skip next instruction if Vx = Vy.
		if CHIP8.Cpu.v[(b&0x0F00)>>8] == CHIP8.Cpu.v[(b&0x00F0)>>4] {
			CHIP8.Cpu.Pc += 2
		}
	case 0x6000:
		//0x6XNN LD Vx, byte -> Set Vx = kk.
		CHIP8.Cpu.v[(b&0x0F00)>>8] = uint8(b & 0x00FF)
	case 0x7000:
		//0x7XNN ADD Vx, byte -> Set Vx = Vx + kk.
		CHIP8.Cpu.v[(b&0x0F00)>>8] += uint8(b & 0x00FF)
	case 0x8000:
		switch b & 0x000F {
		case 0x0000:
			//0x8XY0 LD Vx, Vy -> Set Vx = Vy.
			CHIP8.Cpu.v[(b&0x0F00)>>8] = CHIP8.Cpu.v[(b&0x00F0)>>4]
		case 0x0001:
			//0x8XY1 OR Vx, Vy -> Set Vx = Vx OR Vy.
			CHIP8.Cpu.v[(b&0x0F00)>>8] |= CHIP8.Cpu.v[(b&0x00F0)>>4]
		case 0x0002:
			//0x8XY2 AND Vx, Vy -> Set Vx = Vx AND Vy.
			CHIP8.Cpu.v[(b&0x0F00)>>8] &= CHIP8.Cpu.v[(b&0x00F0)>>4]
		case 0x0003:
			//0x8XY3 XOR Vx, Vy -> Set Vx = Vx XOR Vy.
			CHIP8.Cpu.v[(b&0x0F00)>>8] ^= CHIP8.Cpu.v[(b&0x00F0)>>4]
		case 0x0004:
			//0x8XY4 ADD Vx, Vy -> Set Vx = Vx + Vy, set VF = carry.
			if CHIP8.Cpu.v[(b&0x00F0)>>4] > (0xFF - CHIP8.Cpu.v[(b&0x0F00)>>8]) {
				CHIP8.Cpu.v[0xF] = 1
			} else {
				CHIP8.Cpu.v[0xF] = 0
			}
			CHIP8.Cpu.v[(b&0x0F00)>>8] += CHIP8.Cpu.v[(b&0x00F0)>>4]
		case 0x0005:
			//0x8XY5 SUB Vx, Vy -> Set Vx = Vx - Vy, set VF = NOT borrow.
			if CHIP8.Cpu.v[(b&0x00F0)>>4] > CHIP8.Cpu.v[(b&0x0F00)>>8] {
				CHIP8.Cpu.v[0xF] = 0
			} else {
				CHIP8.Cpu.v[0xF] = 1
			}
			CHIP8.Cpu.v[(b&0x0F00)>>8] -= CHIP8.Cpu.v[(b&0x00F0)>>4]
		case 0x0006:
			//0x8XY6 SHR Vx {, Vy} -> Set Vx = Vx SHR 1.
			CHIP8.Cpu.v[0xF] = CHIP8.Cpu.v[(b&0x0F00)>>8] & 0x1
			CHIP8.Cpu.v[(b&0x0F00)>>8] >>= 1
		case 0x0007:
			//0x8XY7 SUBN Vx, Vy -> Set Vx = Vy - Vx, set VF = NOT borrow.
			if CHIP8.Cpu.v[(b&0x0F00)>>8] > CHIP8.Cpu.v[(b&0x00F0)>>4] {
				CHIP8.Cpu.v[0xF] = 0
			} else {
				CHIP8.Cpu.v[0xF] = 1
			}
			CHIP8.Cpu.v[(b&0x0F00)>>8] = CHIP8.Cpu.v[(b&0x00F0)>>4] - CHIP8.Cpu.v[(b&0x0F00)>>8]
		case 0x000E:
			//0x8XYE SHL Vx {, Vy} -> Set Vx = Vx SHL 1.
			CHIP8.Cpu.v[0xF] = CHIP8.Cpu.v[(b&0x0F00)>>8] >> 7
			CHIP8.Cpu.v[(b&0x0F00)>>8] <<= 1
		}
	case 0x9000:
		//0x9XY0 SNE Vx, Vy -> Skip next instruction if Vx != Vy.
		if CHIP8.Cpu.v[(b&0x0F00)>>8] != CHIP8.Cpu.v[(b&0x00F0)>>4] {
			CHIP8.Cpu.Pc += 2
		}
	case 0xA000:
		//0xANNN LD I, addr -> Set I = nnn.
		CHIP8.Cpu.i = b & 0x0FFF
	case 0xB000:
		//0xBNNN JP V0, addr -> Jump to location nnn + V0.
		CHIP8.Cpu.Pc = (b & 0x0FFF) + uint16(CHIP8.Cpu.v[0])
	case 0xC000:
		//0xCXNN RND Vx, byte -> Set Vx = random byte AND nn.
		CHIP8.Cpu.v[(b&0x0F00)>>8] = uint8(rand.Intn(255)) & uint8(b&0x00FF)
	case 0xD000:
		//0xDXYK
		var tmps []uint8
		for i := CHIP8.Cpu.i; i < CHIP8.Cpu.i+b&0x000F; i++ {
			tmps = append(tmps, CHIP8.Cpu.memory[i])
		}
		for i := uint8(0); i < uint8(len(tmps)); i++ {
			for j := uint8(0); j < 8; j++ {
				CHIP8.Screen.Mapscreen[(CHIP8.Cpu.v[(b&0x0F00)>>8]+j)%64][(CHIP8.Cpu.v[(b&0x00F0)>>4]+i)%32] ^= tmps[i] >> (7 - j) & 0x01
			}
		}
	case 0xE000:
		switch b & 0x000F {
		case 0x000E:
			fmt.Println("\nSKNP Vx ----------------------------------------------------------------------------- Not Pressed")
			//0xEX9E SKP Vx -> Skip next instruction if key with the value of Vx is pressed.
			if CHIP8.clavier.isPressed[CHIP8.Cpu.v[(b&0x0F00)>>8]] {
				fmt.Println("\nSKP Vx ----------------------------------------------------------------------------- Pressed")
				CHIP8.Cpu.Pc += 2
			}
		case 0x0001:

			//0xEXA1 SKNP Vx -> Skip next instruction if key with the value of Vx is not pressed.
			if !CHIP8.clavier.isPressed[CHIP8.Cpu.v[(b&0x0F00)>>8]] {
				CHIP8.Cpu.Pc += 2
			}
		}
	case 0xF000:
		switch b & 0x000F {
		case 0x0007:
			//0xFX07 LD Vx, DT -> Set Vx = delay timer value.
			CHIP8.Cpu.v[(b&0x0F00)>>8] = CHIP8.Cpu.dt
		case 0x000A:
			//0xFX0A LD Vx, K -> Wait for a key press, store the value of the key in Vx.
			CHIP8.Cpu.v[(b&0x0F00)>>8] = CHIP8.Cpu.dt
		case 0x0005:
			switch b & 0x00F0 {
			case 0x0010:
				//0xFX15 LD DT, Vx -> Set delay timer = Vx.
				CHIP8.Cpu.dt = CHIP8.Cpu.v[(b&0x0F00)>>8]
			case 0x0050:
				//0xFX55 LD [I], Vx -> Store registers V0 through Vx in memory starting at location I.
				for i := uint16(0); i <= (b&0x0F00)>>8; i++ {
					CHIP8.Cpu.memory[CHIP8.Cpu.i+i] = CHIP8.Cpu.v[i]
				}
			case 0x0060:
				//0xFX65 LD Vx, [I] -> Read registers V0 through Vx from memory starting at location I.
				for i := uint16(0); i <= (b&0x0F00)>>8; i++ {
					CHIP8.Cpu.v[i] = CHIP8.Cpu.memory[CHIP8.Cpu.i+i]
				}
			}
		case 0x0008:
			//0xFX18 LD ST, Vx -> Set sound timer = Vx.
			CHIP8.Cpu.st = CHIP8.Cpu.v[(b&0x0F00)>>8]
		case 0x000E:
			//0xFX1E ADD I, Vx -> Set I = I + Vx.
			CHIP8.Cpu.i += uint16(CHIP8.Cpu.v[(b&0x0F00)>>8])
		case 0x0009:
			//0xFX29 LD F, Vx -> Set I = location of sprite for digit Vx.
			CHIP8.Cpu.i = uint16(CHIP8.Cpu.v[(b&0x0F00)>>8] * 5)
		case 0x0003:
			//0xFX33 LD B, Vx -> Store BCD representation of Vx in memory locations I, I+1, and I+2.
			CHIP8.Cpu.memory[CHIP8.Cpu.i] = CHIP8.Cpu.v[(b&0x0F00)>>8] / 100
			CHIP8.Cpu.memory[CHIP8.Cpu.i+1] = (CHIP8.Cpu.v[(b&0x0F00)>>8] / 10) % 10
			CHIP8.Cpu.memory[CHIP8.Cpu.i+2] = (CHIP8.Cpu.v[(b&0x0F00)>>8] % 100) % 10
		}
	}
}
