package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/walonCode/readmeMaker/cmd"
	"github.com/walonCode/readmeMaker/internal/types"
	"github.com/walonCode/readmeMaker/internal/utils"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading the .env file")
	}

	var projectName, model, license string
	var  contribute bool

	flag.StringVar(&projectName, "projectName", "", "This is the name of the repo")
	flag.StringVar(&model, "model", "mistral", "This is the name of the ai you want to use")
	flag.StringVar(&license, "license", "", "The type of license you want to generate")
	flag.BoolVar(&contribute, "contribute", false, "weather you want to generate a contribute file")

	flag.Parse()

	resultChan := make(chan types.GenResult)
	var wg sync.WaitGroup

	if license != "" {
		wg.Add(1)
		go cmd.License(model, license, resultChan, &wg)
	}

	if contribute {
		wg.Add(1)
		go cmd.Contribute(model, projectName, resultChan, &wg)
	}

	if projectName != "" {
		wg.Add(1)
		go cmd.Readme(projectName, model, resultChan, &wg)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		if result.Err != nil {
			log.Printf("❌ Failed to generate %s: %v\n", result.Filename, result.Err)
		} else {
			err := utils.WriteFile(result.Content, result.Filename)
			if err != nil {
				log.Printf("❌ Failed to write %s: %v\n", result.Filename, err)
			} else {
				log.Printf("✅ %s generated successfully\n", result.Filename)
			}
		}
	}

}
