package matasano

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
)

func HexToBase64(s string) (string, error) {
	bytes, err := hex.DecodeString(s)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
}

func FixedXOR(a, b string) (string, error) {
	decoded_a, err_a := hex.DecodeString(a)
	decoded_b, err_b := hex.DecodeString(b)

	if err_a != nil {
		return "", err_a
	}
	if err_b != nil {
		return "", err_b
	}
	if len(decoded_a) != len(decoded_b) {
		return "", errors.New("The given strings are not of the same length")
	}

	xor := make([]byte, len(decoded_a))

	for i := range decoded_a {
		xor[i] = decoded_a[i] ^ decoded_b[i]
	}

	return hex.EncodeToString(xor), nil
}

func SingleByteXOR(s string, b byte) (string, error) {
	decoded, err := hex.DecodeString(s)
	xor := make([]byte, len(decoded))

	for i, c := range decoded {
		xor[i] = c ^ b
	}

	return string(xor), err
}
