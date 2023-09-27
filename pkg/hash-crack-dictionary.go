package pkg

import (
	"basic-password-breaker/internal"
	"fmt"
)

func CrackDict(passwordMD5 string, filePath string) (string, error) {
	wordlist, err := internal.GetWordlist(filePath)
	if err != nil {
		return "", err
	}

	for _, w := range wordlist {
		wordMD5 := internal.ToMD5(w)
		if passwordMD5 == wordMD5 {
			return w, nil
		}
	}

	return "", &internal.CustomError{
		Message: fmt.Sprintf("Password does not found in the list: %s", filePath),
	}
}
