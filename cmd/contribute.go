package cmd

import (
	"sync"

	"github.com/walonCode/readmeMaker/internal/llm"
	"github.com/walonCode/readmeMaker/internal/types"
)

func Contribute(model string, resultChan chan <- types.GenResult, wg *sync.WaitGroup) {
	defer wg.Done()

	prompt := `Generate a CONTRIBUTING.md file for an open-source project.

Only include:
- Introduction encouraging contributions
- How to fork, clone, and set up the project
- Branch naming convention
- Guidelines for submitting issues and pull requests
- Code style or linting rules (if any)
- Communication etiquette (brief)

Do not include any extra AI commentary or headers. Just return the full content of CONTRIBUTING.md ready to be saved to a file.

`
	content, err := llm.AiGeneration(prompt, model)

	resultChan <- types.GenResult{
		Filename: "Contribute.md",
		Content: content,
		Err:err,
	}
	
}