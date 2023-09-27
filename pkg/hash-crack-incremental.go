package pkg

import (
	"basic-password-breaker/internal"
	"fmt"
	"strings"
)

func CrackIncr(md5 string, passwordLength int, currentPass []string) (string, error) {
	alphabet := internal.GetAlphabet()

	if passwordLength >= 1 {
		if len(currentPass) == 0 {
			c := strings.Repeat("a", passwordLength)
			cs := strings.Split(c, "")
			return CrackIncr(md5, passwordLength, cs)
		}

		for _, letter := range alphabet {
			currentPass[passwordLength-1] = letter
			cstr := strings.Join(currentPass, "")
			fmt.Printf("[TENTATIVE] %s\n", cstr)
			currentMD5 := internal.ToMD5(cstr)
			if currentMD5 == md5 {
				return cstr, nil
			}

			res, _ := CrackIncr(md5, passwordLength-1, currentPass)
			if res != "" {
				return res, nil
			}
		}
	}

	return "", &internal.CustomError{
		Message: "Password does not found: size different than expected length",
	}
}
