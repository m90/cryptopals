package cryptopals

import (
	"testing"
)

func TestHexToBase64(t *testing.T) {
	result, err := HexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if result != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
		t.Errorf("Unexpected result %s", result)
	}
}

func TestFixedXOR(t *testing.T) {
	result, err := FixedXOR("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if result != "746865206b696420646f6e277420706c6179" {
		t.Errorf("Unexpected result %s", result)
	}
}

func TestDeXOR(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	result, err := DeXOR(input, []byte("X"))
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if result != "Cooking MC's like a pound of bacon" {
		t.Errorf("Unexpected result %s", result)
	}
}

func TestRepeatingKeyXOR(t *testing.T) {
	input := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	expected := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	key := "ICE"
	result, err := RepeatingKeyXOR(input, []byte(key))
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if result != expected {
		t.Errorf("Unexpected result %v", result)
	}
}
