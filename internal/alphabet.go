package internal

import (
	"unicode"
)

func GetAlphabet() [52]string {
	var alphabet [52]string
	for i := 0; i < 26; i++ {
		minl := rune('a' + i)
		majl := unicode.ToUpper(minl)
		alphabet[i] = string(minl)
		alphabet[26+i] = string(majl)
	}
	return alphabet
}
