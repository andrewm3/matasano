// Package matasano contains utilities to solve the Matasano Crypto Challenges
// This file covers the basics.
package matasano

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"strings"
)

// HexToBase64 converts a hex encoded string to a base64 encoded one.
func HexToBase64(s string) (string, error) {
	bytes, err := hex.DecodeString(s)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
}

// FixedXOR returns the XOR of two strings of the same length.
func FixedXOR(a, b string) (string, error) {
	decodedA, errA := hex.DecodeString(a)
	decodedB, errB := hex.DecodeString(b)

	if errA != nil {
		return "", errA
	}
	if errB != nil {
		return "", errB
	}
	if len(decodedA) != len(decodedB) {
		return "", errors.New("The given strings are not of the same length")
	}

	xor := make([]byte, len(decodedA))

	for i := range decodedA {
		xor[i] = decodedA[i] ^ decodedB[i]
	}

	return hex.EncodeToString(xor), nil
}

// SingleByteXOR takes a string and XOR's it against a single character.
func SingleByteXOR(s string, b byte) (string, error) {
	decoded, err := hex.DecodeString(s)
	xor := make([]byte, len(decoded))

	for i, c := range decoded {
		xor[i] = c ^ b
	}

	return string(xor), err
}

// RepeatingKeyXOR takes a string and XOR's it against a repeating key.
func RepeatingKeyXOR(s []byte, key []byte) []byte {
	xor := make([]byte, len(s))

	var j int

	for i, c := range s {
		xor[i] = c ^ key[j]
		if j++; j >= len(key) {
			j = 0
		}
	}

	return xor
}

// Englishness calculates the likelihood that the given text is an English phrase.
func Englishness(s string) float64 {
	// Frequencies sourced from https://en.wikipedia.org/wiki/Letter_frequency
	charFrequencies := map[rune]float64{
		'a': 0.08167, 'b': 0.01492, 'c': 0.02782, 'd': 0.04253, 'e': 0.12702, 'f': 0.02228,
		'g': 0.02015, 'h': 0.06094, 'i': 0.06966, 'j': 0.00153, 'k': 0.00772, 'l': 0.04025,
		'm': 0.02406, 'n': 0.06749, 'o': 0.07507, 'p': 0.01929, 'q': 0.00095, 'r': 0.05987,
		's': 0.06327, 't': 0.09056, 'u': 0.02758, 'v': 0.00978, 'w': 0.02360, 'x': 0.00150,
		'y': 0.01974, 'z': 0.00074, ' ': 0.15000,
	}

	sum := 0.0

	for _, r := range strings.ToLower(s) {
		if f, ok := charFrequencies[r]; ok {
			sum += f
		}
	}

	return sum
}
