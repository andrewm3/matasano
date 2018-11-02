// Package matasano contains utilities to solve the Matasano Crypto Challenges
// This file covers the basics.
package matasano

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
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
