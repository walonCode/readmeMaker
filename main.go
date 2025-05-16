package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	fileparser "github.com/walonCode/readmeMaker/internal/file_parser"
	"github.com/walonCode/readmeMaker/internal/llm"
	"github.com/walonCode/readmeMaker/internal/tree"
	"github.com/walonCode/readmeMaker/internal/utils"
)

const llama = "llama3-8b-8192"
const mistral = "mistral-saba-24b"
const qwen = "qwen-qwq-32b"

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	var model,projectName string
	flag.StringVar(&projectName, "projectName","", "The project Name")
	flag.StringVar(&model, "model", "llama", "The ai you want to make the readme")

	flag.Parse()

	if projectName == "" {
		fmt.Println("❌ Error: --projectName name is required")
		flag.Usage()
		os.Exit(1)
	}

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

	prompt, err := utils.BuildPrompt("readmeTree.txt", projectName)
	if err != nil {
		fmt.Println("error",err)
		return
	}

	f.Close()

	var readme string
	
	if model == "llama" {
		aiResponse, err := llm.GenerateReadme(prompt, llama)
		if err != nil {
			fmt.Println("❌ Failed to generate readme with llama Ai model ")
		}
		readme = aiResponse
	}else if model == "mistral" {
		aiResponse, err := llm.GenerateReadme(prompt, mistral)
		if err != nil {
			fmt.Println("❌ Failed to generate readme with mistral Ai model ",err)
		}

		readme = aiResponse
	}else if model == "qwen" {
		aiResponse, err := llm.GenerateReadme(prompt, qwen)
		if err != nil {
			fmt.Println("❌ Failed to generate readme with qwen Ai model ")
		}

		readme = aiResponse
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