package main

import (
	"basic-password-breaker/internal"
	"basic-password-breaker/pkg"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

func main() {
	password, err := internal.GetInput()
	if err != nil {
		log.Fatalln("Cannot scan the password write previously ...")
	}

	var result string
	var currentLetter string
	alphabet := pkg.GetAlphabet()
	for _, letter := range strings.Split(password, "") {
		for letter != currentLetter {
			fmt.Println(result + currentLetter)
			time.Sleep(5 * time.Millisecond)

			randomIdx := rand.Intn(len(alphabet))
			currentLetter = alphabet[randomIdx]
		}
		result += currentLetter
	}
	fmt.Println(result)
}
