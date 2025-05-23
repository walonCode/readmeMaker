package utils

import "os"

func WriteFile(content,filename string) error {
	return os.WriteFile(filename, []byte(content),0664)
}