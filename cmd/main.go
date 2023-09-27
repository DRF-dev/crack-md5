package main

import (
	"basic-password-breaker/internal"
	"basic-password-breaker/pkg"
	"fmt"
	"github.com/gookit/color"
	"os"
	"time"
)

func getCrackingMethod(md5 string, dictionaryPath string, passwordLength int) (string, error) {
	if dictionaryPath != "" {
		return pkg.CrackDict(md5, dictionaryPath)
	}
	return pkg.CrackIncr(md5, passwordLength, []string{})
}

func main() {
	start := time.Now()

	dictionaryPath, password, md5, passwordLength := internal.GetFlags()
	if md5 == "" {
		md5 = internal.ToMD5(password)
		color.Greenf("[MD5 HASH]: %s (%s)\n", md5, password)
		os.Exit(0)
	}

	result, err := getCrackingMethod(md5, dictionaryPath, passwordLength)
	if err != nil {
		color.Redln(err)
		os.Exit(1)
	}
	color.Greenf("Password found with success: %s (%s)\n", result, md5)

	duration := time.Since(start)
	fmt.Printf("Execution time : %s\n", duration)
}
