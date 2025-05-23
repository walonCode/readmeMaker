package cmd

import (
	"flag"
	"fmt"
	"os"
	"sync"

	fileparser "github.com/walonCode/readmeMaker/internal/file_parser"
	"github.com/walonCode/readmeMaker/internal/llm"
	"github.com/walonCode/readmeMaker/internal/tree"
	"github.com/walonCode/readmeMaker/internal/utils"
	"github.com/walonCode/readmeMaker/internal/types"
)



func Readme(projectName, model string, resultChan chan <- types.GenResult, wg *sync.WaitGroup ) {
	defer wg.Done()

	
	if projectName == "" {
		fmt.Println("❌ Error: --projectName name is required")
		flag.Usage()
		os.Exit(1)
	}


	tb := &tree.TreeBuilder{
		Root: ".",
	}

	err := tb.WalkAndCount()
	if err != nil {
		fmt.Println("Error creating the tree: ",err)
		return
	}

	fmt.Printf("📄 Files: %d | 📁 Folders: %d\n", tb.FileCount, tb.FolderCount)

	err = tb.WriteTreeToFile("readmeTree.txt")
	if err != nil {
		fmt.Println("Error creating the readmeTree.txt: ",err)
		return
	}

	snippet, err := fileparser.ParseFile(".")
	if err != nil {
		fmt.Println("Error parsing the file: ",err)
		return
	}

	f, err := os.OpenFile("readmeTree.txt", os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println("Error open the readmeTree.txt: ",err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(snippet); err != nil {
		fmt.Println("Error creating the snippet: ",err)
		return
	}

	fmt.Println("✅ saved to tree.txt")

	prompt, err := utils.BuildPrompt("readmeTree.txt", projectName)
	if err != nil {
		fmt.Println("Error creating the prompt: ",err)
		return
	}

	f.Close()

	
	aiResponse, err := llm.AiGeneration(prompt, model)
	if err != nil {
		fmt.Println("Ai generation error: ",err)
		return
	}

	err = os.Remove("readmeTree.txt")
	if err != nil {
		fmt.Println("Error deleting the readmeTree.txt file: ",err)
		return
	}

	resultChan <- types.GenResult{
		Filename: "README.md",
		Content: aiResponse,
		Err: err,
	}
}