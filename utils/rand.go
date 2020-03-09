package utils

import (
	"math/rand"
	"strings"
	"time"
)

const (
	RandStrModeNumber      = 1
	RandStrModeLetterLower = 2
	RandStrModeLetterUpper = 4
)

func RandStr(length, mode int) string {

	var (
		numbers      = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
		lowerLetters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
		upperLetters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	)
	var chars []string
	var source []string

	if (mode & RandStrModeNumber) == RandStrModeNumber {
		source = append(source, numbers...)
	}

	if (mode & RandStrModeLetterLower) == RandStrModeLetterLower {
		source = append(source, lowerLetters...)
	}

	if (mode & RandStrModeLetterUpper) == RandStrModeLetterUpper {
		source = append(source, upperLetters...)
	}

	sourceLength := len(source)

	for i := 0; i < length; i++ {

		chars = append(chars, source[rand.Intn(sourceLength-1)])
	}

	return strings.Join(chars, "")
}

func init() {
	rand.Seed(time.Now().Unix())
}
