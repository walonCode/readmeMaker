package tree

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var ignoredFolders = map[string]bool{
	"node_modules":true,
	".git" : true,
	"dist": true,
	".next" : true,
	".env":true,
}

type TreeBuilder struct {
	Builder     strings.Builder
	FileCount   int
	FolderCount int
	Root        string
}

func (tb *TreeBuilder)WalkAndCount() error{
	tb.Builder.WriteString("📁 " + filepath.Base(tb.Root) + "\n")

	err := filepath.WalkDir(tb.Root, func(path string, d os.DirEntry, err error)error {
		if err != nil {
			return err
		}

		if path == tb.Root {
			return nil
		}

		if d.IsDir() && ignoredFolders[d.Name()] {
			return filepath.SkipDir
		}
		
		relPath,_ := filepath.Rel(tb.Root, path)
		depth := strings.Count(relPath, string(os.PathSeparator))
		prefix := strings.Repeat("|  ", depth)

		connector := "├── "

		tb.Builder.WriteString(fmt.Sprintf("%s%s%s\n", prefix, connector, d.Name()))

		if d.IsDir() {
			tb.FolderCount++
		} else {
			tb.FileCount++
		}

		return nil
	})

	return err
}

func (tb *TreeBuilder) WriteTreeToFile(filename string)error{
	return os.WriteFile(filename, []byte(tb.Builder.String()), 0664)
}