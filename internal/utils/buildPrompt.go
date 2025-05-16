package utils

import "os"

func BuildPrompt(treePath, projectName string) (string, error) {
	data, err := os.ReadFile(treePath)
	if err != nil {
		return "",err
	}
	prompt := `
		You are an expert documentation writer.
		Using the following project structure and code snippets, write a professional  README.md that is written like a senior software engineer.

		Make sure to include:
		- Project name (guess from structure if not stated)
		- Description
		- Features
		- Installation
		- Usage
		- Technologies
		- Configuration and Env (if an env file is provided)
		- Folder structure (nicely formatted)
		- Authors name and github handle/X handle used dummy text
		- Contribution instructions if appropriate

		Note: Remove all unwanted text from the final response i want a clean readme, also read the snippet to get code concept better thank you

		--- BEGIN STRUCTURE + SNIPPETS ---
		` + string(data) + projectName + "\n--- END ---"

	return prompt, nil
}