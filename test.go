package main

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"io"
	"math"
	"time"
)

func main() {
	// Initialize the audio context
	context := audio.NewContext(44100)

	// Create a player with a sinusoidal wave as the source
	player, err := context.NewPlayer(sineWave(440, 1)) // 440 Hz sine wave for 1 second
	if err != nil {
		fmt.Println("Error creating player:", err)
		return
	}

	// Play the sound
	fmt.Printf("\n", player)
	player.Play()

	//wait
	time.Sleep(time.Second * 1)
	fmt.Printf("\n", player)

	player.Pause()

	// Close the context when the function returns
	player.Close()
}

// sineWave generates a sinusoidal wave with the given frequency and duration
func sineWave(freq float64, durationSeconds int) io.Reader {
	const sampleRate = 44100
	numSamples := sampleRate * durationSeconds
	samples := make([]int, numSamples)

	for i := 0; i < numSamples; i++ {
		samples[i] = int(math.MaxInt16 * math.Sin(2*math.Pi*freq*float64(i)/sampleRate))
	}

	return bytes.NewReader(int16SliceToBytes(samples))
}

// int16SliceToBytes converts a slice of int16 to a byte slice
func int16SliceToBytes(samples []int) []byte {
	byteSlice := make([]byte, len(samples)*2)
	for i, sample := range samples {
		byteSlice[i*2] = byte(sample)
		byteSlice[i*2+1] = byte(sample >> 8)
	}
	return byteSlice
}
