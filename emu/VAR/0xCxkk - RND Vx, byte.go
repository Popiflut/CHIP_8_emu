package VAR

import "math/rand"

// 0xCXNN RND Vx, byte -> Set Vx = random byte AND nn.
func RND_Vx_byte(b uint16) {
	CHIP8.Cpu.V[(b&0x0F00)>>8] = uint8(rand.Intn(255)) & uint8(b&0x00FF)
}
