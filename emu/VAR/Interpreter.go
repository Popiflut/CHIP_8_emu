package VAR

func (cpu *CPUs) Interpreter(b uint16) {
	switch b & 0xF000 {
	case 0x0000:
		switch b & 0x000F {
		case 0x0000:
			CLS()
		case 0x000E:
			RET()
		}
	case 0x1000:
		JP_addr(b)
	case 0x2000:
		CALL_addr(b)
	case 0x3000:
		SE_Vx_byte(b)
	case 0x4000:
		SNE_Vx_byte(b)
	case 0x5000:
		SE_Vx_Vy(b)
	case 0x6000:
		LD_Vx_byte(b)
	case 0x7000:
		ADD_Vx_byte(b)
	case 0x8000:
		switch b & 0x000F {
		case 0x0000:
			LD_Vx_Vy(b)
		case 0x0001:
			OR_Vx_Vy(b)
		case 0x0002:
			AND_Vx_Vy(b)
		case 0x0003:
			XOR_Vx_Vy(b)
		case 0x0004:
			ADD_Vx_Vy(b)
		case 0x0005:
			SUB_Vx_Vy(b)
		case 0x0006:
			SHR_Vx_Vy(b)
		case 0x0007:
			SUBN_Vx_Vy(b)
		case 0x000E:
			SHL_Vx_Vy(b)
		}
	case 0x9000:
		SNE_Vx_Vy(b)
	case 0xA000:
		LD_I_addr(b)
	case 0xB000:
		JP_V0_addr(b)
	case 0xC000:
		RND_Vx_byte(b)
	case 0xD000:
		DRW_Vx_Vy_nibble(b)
	case 0xE000:
		switch b & 0x000F {
		case 0x000E:
			SKP_Vx(b)
		case 0x0001:
			SKNP_Vx(b)
		}
	case 0xF000:
		switch b & 0x000F {
		case 0x0007:
			LD_Vx_DT(b)
		case 0x000A:
			LD_Vx_K(b)
		case 0x0005:
			switch b & 0x00F0 {
			case 0x0010:
				LD_DT_Vx(b)
			case 0x0050:
				LD_I_Vx(b)
			case 0x0060:
				LD_Vx_I(b)
			}
		case 0x0008:
			LD_ST_Vx(b)
		case 0x000E:
			ADD_I_Vx(b)
		case 0x0009:
			LD_F_Vx(b)
		case 0x0003:
			LD_B_Vx(b)
		}
	}
}
