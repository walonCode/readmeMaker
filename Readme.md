# ReadmeMaker

## Description
readmeMaker is a Go-based tool designed to automate the generation of README files, contribution guidelines, and license files using Natural Language Models (LLMs). The tool integrates with an AI service to generate high-quality documentation based on predefined templates and input parameters.

## Features
- Generate README files from specified project information.
- Create contribution guidelines using AI.
- Generate license files based on the specified licenses.
- Supports multiple AI models for generation.
- Parallel processing for faster execution.

## Installation
Ensure you have Go installed on your system. Clone the repository and install the necessary dependencies:

```sh
git clone https://github.com/walonCode/readmeMaker.git
cd readmeMaker
go mod tidy
```

## Usage
Configure your environment variables in the `.env` file:

```plaintext
API_KEY="your_api_key_here"
```

Run the tool using the command line with the required parameters:

```sh
go run main.go -projectName "your_project_name" -model "your_ai_model" -license -contribute
```

- `-projectName`: Specifies the name of the project.
- `-model`: Specifies the AI model to use for generation[mistral,gema,llama].
- `-license`: Option to generate a license file.
- `-contribute`: Option to generate a contribution guidelines file.

## Technologies
- **Go**: Main programming language.
- **godotenv**: Library for loading environment variables from a .env file.

## Configuration and Environment Variables
Create a `.env` file in the root directory of the project with the following content:

```plaintext
API_KEY="your_api_key_here"
```

## Folder Structure
```plaintext
.
├── .env
├── .env.example
├── .gitignore
├── cmd
│   ├── contribute.go
│   ├── license.go
│   ├── readme.go
├── go.mod
├── go.sum
├── internal
│   ├── file_parser
│   │   ├── fileParser.go
│   ├── llm
│   │   ├── ai.services.go
│   ├── tree
│   │   ├── buildTree.go
│   ├── types
│   │   ├── types.go
│   ├── utils
│   │   ├── buildPrompt.go
│   │   ├── writeReadme.go
└── main.go
```

## Author
- **Name**: Mohamed Lamin Walon-Jalloh
- **GitHub**: @walonCode

## Contribution
Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Commit your changes.
4. Push the branch to your fork.
5. Open a Pull Request.

In your PR, please include a clear description of the changes and why they are necessary. Make sure to update any relevant documentation as well.

```
