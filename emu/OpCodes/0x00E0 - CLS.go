package OpCodes

import "main/emu"

// 0x00E0 CLS -> Clear the display.
func CLS() {
	for i := 0; i < emu.ResolWidth; i++ {
		for j := 0; j < emu.ResolHeight; j++ {
			emu.CHIP8.Screen.Mapscreen[i][j] = 0
		}
	}
}
