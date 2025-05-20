package fileparser

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var targetFiles = map[string]bool{
	"package.json":true,
	".env.example":true,
	"go.mod":true,
	"main.go":true,
}

func readLines(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "",err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var result strings.Builder


	for scanner.Scan(){
		line := scanner.Text()

		result.WriteString(line + "\n")
	}

	if err := scanner.Err(); err != nil {
		return "",err
	}

	return result.String(),nil
}

func ParseFile(root string)(string, error) {
	var snippetBuilder strings.Builder
	snippetBuilder.WriteString("\n\n\n🧩 Snippet from key Files: \n")

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && targetFiles[d.Name()]{
			snippetBuilder.WriteString(fmt.Sprintf("\n🔹 %s\n", d.Name()))

			if targetFiles[d.Name()] && filepath.Dir(path) != root {
				return nil
			}
			
			content, err := readLines(path)
			if err != nil {
				snippetBuilder.WriteString(" [Failed to read snippet] \n")
			}else {
				snippetBuilder.WriteString("```\n" + content + "```\n")
			}
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	return snippetBuilder.String(), nil
}