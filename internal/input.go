package internal

import (
	"fmt"
)

func GetInput() (password string, err error) {
	fmt.Print("What is the password to crack :")
	_, err = fmt.Scan(&password)
	if err != nil {
		return password, err
	}
	return password, nil
}
