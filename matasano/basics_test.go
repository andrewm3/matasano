package matasano

import (
	"encoding/hex"
	"os"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	should := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	result, err := HexToBase64(input)

	if err != nil {
		t.Error("Error occurred during conversion")
	}
	if result != should {
		t.Error("Expected", should, "not", result)
	}
}

func TestFixedXOR(t *testing.T) {
	inputA := "1c0111001f010100061a024b53535009181c"
	inputB := "686974207468652062756c6c277320657965"
	should := "746865206b696420646f6e277420706c6179"

	result, err := FixedXOR(inputA, inputB)

	if err != nil {
		t.Error("Error occurred during function")
	}
	if result != should {
		t.Error("Expected", should, "not", result)
	}
}

func TestSingleByteXOR(t *testing.T) {
	var key byte = 'X'
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	should := "Cooking MC's like a pound of bacon"

	result, err := SingleByteXOR(input, key)

	if err != nil {
		t.Error("Error occurred during function")
	}
	if result != should {
		t.Error("Expected", should, "not", result)
	}
}

func TestRepeatingKeyXOR(t *testing.T) {
	input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	should := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272" +
		"a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	key := "ICE"

	result := hex.EncodeToString(RepeatingKeyXOR([]byte(input), []byte(key)))

	if result != should {
		t.Error("Expected", should, "not", result)
	}
}

func TestEnglishness(t *testing.T) {
}

func TestDecryptSingleByteXOR(t *testing.T) {
	hex := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	should := "Cooking MC's like a pound of bacon"
	decrypted, err := DecryptSingleByteXOR(hex)
	if err != nil {
		t.Error("error occurred during function")
	}
	if decrypted.phrase != should {
		t.Error("Expected", should, "not", decrypted.phrase)
	}
}

func BenchmarkDecryptSingleByteXOR(b *testing.B) {
	hex := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	for n := 0; n < b.N; n++ {
		DecryptSingleByteXOR(hex)
	}
}

func TestDetectSingleByteXOR(t *testing.T) {
	file, err := os.Open("../resources/set-1-challenge-4.txt")
	defer file.Close()
	should := "Now that the party is jumping\n"
	result, err := DetectSingleByteXOR(file)
	if err != nil {
		t.Error("error occurred during function")
	}
	if result.phrase != should {
		t.Error("Expected", should, "not", result.phrase)
	}
}

func TestHammingDistance(t *testing.T) {
	should := 37
	result, err := HammingDistance("this is a test", "wokka wokka!!!")
	if err != nil {
		t.Error("error occurred during function")
	}
	if result != should {
		t.Error("expected", should, "not", result)
	}
}
