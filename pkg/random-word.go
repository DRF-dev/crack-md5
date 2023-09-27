package pkg

import (
	"basic-password-breaker/internal"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func RandomWord(password string) string {
	var result string
	var currentLetter string

	alphabet := internal.GetAlphabet()

	for _, letter := range strings.Split(password, "") {
		for letter != currentLetter {
			fmt.Println(result + currentLetter)
			time.Sleep(5 * time.Millisecond)

			randomIdx := rand.Intn(len(alphabet))
			currentLetter = alphabet[randomIdx]
		}
		result += currentLetter
	}
	return result
}
