package main

import (
	"fmt"
	"os"

	fileparser "github.com/walonCode/readmeMaker/internal/file_parser"
	"github.com/walonCode/readmeMaker/internal/llm"
	"github.com/walonCode/readmeMaker/internal/tree"
	"github.com/walonCode/readmeMaker/internal/utils"
)


func main() {
	projectDir := "."

	tb := &tree.TreeBuilder{
		Root: projectDir,
	}

	err := tb.WalkAndCount()
	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Printf("📄 Files: %d | 📁 Folders: %d\n", tb.FileCount, tb.FolderCount)

	err = tb.WriteTreeToFile("readmeTree.txt")
	if err != nil {
		fmt.Println("error",err)
		return
	}

	snippet, err := fileparser.ParseFile(projectDir)
	if err != nil {
		fmt.Println("error", err)
	}

	f, err := os.OpenFile("readmeTree.txt", os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println("error",err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(snippet); err != nil {
		fmt.Println("append failed",err)
		return
	}

	fmt.Println("✅ saved to tree.txt")

	prompt, err := utils.BuildPrompt("readmeTree.txt", "Youtube Clone")
	if err != nil {
		fmt.Println("error",err)
		return
	}

	f.Close()
	
	readme, err := llm.GenerateReadme(prompt,"")
	if err != nil {
		fmt.Println("ERROR",err)
		return
	}

	err = utils.WriteReadme(readme)
	if err != nil {
		fmt.Println("error",err)
		return
	}

	fmt.Println("✅ Readme file  created")

	err = os.Remove("readmeTree.txt")

	if err != nil {
		fmt.Println("Failing in deleting file",err)
		return
	}

	fmt.Println("🗑️ created readmeTree.txt delete")


}