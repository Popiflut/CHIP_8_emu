package emu

import (
	"bytes"
	"fmt"
	"io"
	"main/emu/VAR"
	"math"
	"time"
)

func MakeSound() {
	// Create a Player with a sinusoidal wave as the source
	err := error(nil)
	VAR.Player, err = VAR.Context.NewPlayer(sineWave(float64(VAR.Freq), VAR.DurationSeconds)) // 440 Hz sine wave for 1 second
	if err != nil {
		fmt.Println("Error creating Player:", err)
		return
	}

	// Play the sound
	VAR.Player.Play()

	time.Sleep(time.Millisecond)
	// Close the context when the function returns
	VAR.Player.Close()
}

func sineWave(freq float64, durationSeconds int) io.Reader {
	const sampleRate = 44100
	numSamples := sampleRate * durationSeconds
	samples := make([]int, numSamples)

	for i := 0; i < numSamples; i++ {
		samples[i] = int(math.MaxInt16 * math.Sin(2*math.Pi*freq*float64(i)/sampleRate))
	}

	return bytes.NewReader(int16SliceToBytes(samples))
}

func int16SliceToBytes(samples []int) []byte {
	byteSlice := make([]byte, len(samples)*2)
	for i, sample := range samples {
		byteSlice[i*2] = byte(sample)
		byteSlice[i*2+1] = byte(sample >> 8)
	}
	return byteSlice
}
