package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	fileparser "github.com/walonCode/readmeMaker/internal/file_parser"
	"github.com/walonCode/readmeMaker/internal/llm"
	"github.com/walonCode/readmeMaker/internal/types"
)

func Contribute(model, projectName string, resultChan chan <- types.GenResult, wg *sync.WaitGroup) {
	defer wg.Done()

	var language string

	err := filepath.WalkDir(".",func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && fileparser.TargetFiles[d.Name()]{
			if d.Name() == "package.json" {
				language = "Javascript"
			} else if d.Name() == "go.mod"{
				language = "Go"
			} else if d.Name() == "Cargo.toml"{
				language = "Rust"
			} else if d.Name() == "pyproject.toml"{
				language = "Python"
			}
		}

		return nil
	});

	fmt.Println(language)

	if err != nil {
		log.Fatal("error occured when figuring out the language used",err)
	}

	prompt := fmt.Sprintf("Generate a CONTRIBUTING.md file for an open-source project.Only include:- ProjectName:  %v- Introduction encouraging contributions- How to fork, clone, and set up the project- Branch naming convention,- Guidelines for submitting issues and pull requests,- Code style or linting rules of the language used %v,- Communication etiquette (brief),Do not include any extra AI commentary or headers. Just return the full content of CONTRIBUTING.md ready to be saved to a file.using the language used and project name are stated below.Use icon to make it appealing and better looking be explict but concise and generate it as if a senior developer with many years of experience did it.",projectName,language)


	content, err := llm.AiGeneration(prompt, model)

	resultChan <- types.GenResult{
		Filename: "Contribute.md",
		Content: content,
		Err:err,
	}
	
}