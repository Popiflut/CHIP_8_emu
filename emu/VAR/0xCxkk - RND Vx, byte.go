package VAR

import "math/rand"

// 0xCXNN RND Vx, byte -> Set Vx = random byte AND nn.
// Set Vx = random byte AND kk. The interpreter generates a random number from 0 to 255, which is then
// ANDed with the value kk. The results are stored in Vx. See instruction 8xy2 for more information on AND.
func RND_Vx_byte(b uint16) {
	CHIP8.Cpu.V[(b&0x0F00)>>8] = uint8(rand.Intn(255)) & uint8(b&0x00FF)
}
