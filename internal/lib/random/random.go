package random

// TO DO: Generate a random with given number of symbols

import (
	"math/rand"
)

func NewRandomString(aliasLength int) string {
	str := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"1234567890")

	output_string := make([]byte, aliasLength)

	for i := range output_string {
		output_string[i] = str[rand.Intn(len(str))]
	}

	return string(output_string)
}
