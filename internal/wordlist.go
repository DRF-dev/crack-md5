package internal

import (
	"os"
	"strings"
)

func GetWordlist(path string) ([]string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return []string{}, err
	}
	strContent := string(content)
	strContent = strings.ReplaceAll(strContent, "\r", "")
	return strings.Split(strContent, "\n"), nil
}
