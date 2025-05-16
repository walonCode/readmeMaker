package utils

import "os"

func WriteReadme(content string) error {
	return os.WriteFile("Readme.md", []byte(content),0664)
}