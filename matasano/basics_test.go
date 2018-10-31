package matasano

import (
	"testing"
)

func TestHexToBase64(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	should := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	result, err := HexToBase64(input)

	if err != nil {
		t.Error("Error occured during conversion")
	}
	if result != should {
		t.Error("Expected ", should, ", not ", result)
	}
}
