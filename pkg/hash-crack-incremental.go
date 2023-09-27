package pkg

import (
	"basic-password-breaker/internal"
	"strings"
)

func GetCombination(passLength int, current []string, stack *[]string) {
	alphabet := internal.GetAlphabet()

	if passLength >= 1 {
		for _, letter := range alphabet {
			current[passLength-1] = letter
			str := strings.Join(current, "")
			*stack = append(*stack, str)
			GetCombination(passLength-1, current, stack)
		}
	}
}

func worker(md5 string, currentPassChan chan string, resultChan chan string) {
	for c := range currentPassChan {
		currentMD5 := internal.ToMD5(c)
		if currentMD5 == md5 {
			resultChan <- c
		}
	}
}

func CrackIncr(md5 string, passLength int) (string, error) {
	c := strings.Repeat("a", passLength)
	cs := strings.Split(c, "")
	stack := make([]string, 0)
	GetCombination(passLength, cs, &stack)

	const nbrWorkers int = 1000
	currentPassChan := make(chan string)
	resultChan := make(chan string)
	defer close(resultChan)

	for i := 0; i < nbrWorkers; i++ {
		go worker(md5, currentPassChan, resultChan)
	}

	for _, s := range stack {
		currentPassChan <- s
	}
	close(currentPassChan)

	return <-resultChan, nil
}
