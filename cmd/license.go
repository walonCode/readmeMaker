package cmd

import (
	"fmt"
	"sync"

	"github.com/walonCode/readmeMaker/internal/llm"
	"github.com/walonCode/readmeMaker/internal/types"
)

func License(model, licenseType string, resultChan chan <- types.GenResult, wg *sync.WaitGroup) {
	defer wg.Done()

	prompt := fmt.Sprintf("Generate a complete %v License file for an open-source software project. Only return the full text of the license with no extra commentary or explanation.Do not include any additional text or notes. Return only the license content. put a place holder for the author",licenseType)
	
	result, err := llm.AiGeneration(prompt, model)

	resultChan <- types.GenResult{
		Filename: "License.txt",
		Content: result,
		Err: err,
	}

}