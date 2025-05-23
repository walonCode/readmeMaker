package cmd

import (
	"sync"

	"github.com/walonCode/readmeMaker/internal/llm"
	"github.com/walonCode/readmeMaker/internal/types"
)

func License(model string, resultChan chan <- types.GenResult, wg *sync.WaitGroup) {
	defer wg.Done()

	prompt := `Generate a complete MIT License file for an open-source software project. Only return the full text of the license with no extra commentary or explanation.

Include:
- The license title (e.g., "MIT License")
- Year: 2025
- Author: Mohamed Jalloh

Do not include any additional text or notes. Return only the license content.
`
	result, err := llm.AiGeneration(prompt, model)

	resultChan <- types.GenResult{
		Filename: "License.txt",
		Content: result,
		Err: err,
	}

}