package cryptopals

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"regexp"
)

func HexToBase64(in string) (string, error) {
	b, err := hex.DecodeString(in)
	if err != nil {
		return "", fmt.Errorf("error decoding input: %w", err)
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func FixedXOR(left, right string) (string, error) {
	leftBytes, err := hex.DecodeString(left)
	if err != nil {
		return "", fmt.Errorf("error decoding input: %w", err)
	}
	rightBytes, err := hex.DecodeString(right)
	if err != nil {
		return "", fmt.Errorf("error decoding input: %w", err)
	}

	var out []byte
	for idx, _ := range leftBytes {
		out = append(out, leftBytes[idx]^rightBytes[idx])
	}
	return hex.EncodeToString(out), nil
}

func DeXOR(input string, key []byte) (string, error) {
	b, err := hex.DecodeString(input)
	if err != nil {
		return "", fmt.Errorf("error decoding input: %w", err)
	}
	var out []byte
	for idx := range b {
		out = append(out, key[0]^b[idx])
	}
	return fmt.Sprintf("%s", out), nil
}

func RepeatingKeyXOR(input string, key []byte) (string, error) {
	var out []byte
	for idx, byte := range bytes.NewBufferString(input).Bytes() {
		out = append(out, key[idx%len(key)]^byte)
	}
	return hex.EncodeToString(out), nil
}

func scoreString(s string) int {
	isChar := regexp.MustCompile("^[aeiourstlmn\\s]$")
	score := 0
	for _, char := range s {
		if isChar.MatchString(string(char)) {
			score++
		}
	}
	return score
}

type Option struct {
	Result    string
	Qualifier interface{}
}

func ScoreOptions(s []Option) Option {
	var winner Option
	var highScore int
	for _, str := range s {
		score := scoreString(str.Result)
		if score > highScore {
			winner, highScore = str, score
		}
	}
	return winner
}
