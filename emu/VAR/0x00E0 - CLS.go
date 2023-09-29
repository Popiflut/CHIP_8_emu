package VAR

// 0x00E0 CLS -> Clear the display.
// Clear the display.
func CLS() {
	for i := 0; i < ResolWidth; i++ {
		for j := 0; j < ResolHeight; j++ {
			CHIP8.Screen.Mapscreen[i][j] = 0
		}
	}
}
