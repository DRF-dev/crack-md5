package internal

import (
	"flag"
	"fmt"
	"os"
)

func GetInput() (password string, err error) {
	fmt.Print("What is the password to crack : ")
	_, err = fmt.Scan(&password)
	if err != nil {
		return password, err
	}
	return password, nil
}

func GetFlags() (string, string, string, int) {
	dictPath := flag.String("f", "", "Path of dictionary file")
	pass := flag.String("g", "", "Generate MD5 hash of a password")
	md5 := flag.String("md5", "", "MD5 hash of a password")
	passLength := flag.Int("l", 0, "Length of password (real not hash)")
	flag.Parse()

	if *pass == "" && *dictPath == "" && *passLength == 0 {
		fmt.Println("Un des deux flag -f ou -pl est attendu")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *pass == "" && *md5 == "" {
		fmt.Println("Un des deux flag -g ou -md5 est attendu")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *pass != "" && *md5 != "" {
		fmt.Println("Un seul flag est attendu entre -g et -md5")
		flag.PrintDefaults()
		os.Exit(1)
	}

	return *dictPath, *pass, *md5, *passLength
}
