package main

import (
	"fmt"
	"github.com/andrewm3/matasano/matasano"
)

func main() {
	hex := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	var bestEnglishness float64
	var bestPhrase string
	var bestKey byte

	for i := 0; i < 256; i++ {
		xor, err := matasano.SingleByteXOR(hex, byte(i))
		if err != nil {
			fmt.Println(err)
		} else {
			englishness := matasano.Englishness(xor)
			if englishness > bestEnglishness {
				bestEnglishness = englishness
				bestPhrase = xor
				bestKey = byte(i)
			}
		}
	}

	fmt.Println(string(bestKey), bestPhrase, bestEnglishness)
}
