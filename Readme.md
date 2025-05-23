# readmeMaker

## Description

`readmeMaker` is a command-line tool designed to generate readme files, license files, and contribution guidelines. It leverages AI models such as Mistral for generating human-like documentation and ensures a professional and clean output.

## Features

- Generate README.md
- Generate Contribute.md
- Generate License.txt
- Allows specification of the AI model to be used
- Supports flag-based configuration options for specifying the content to generate

## Installation

1. Clone the repository from [GitHub](https://github.com/walonCode/readmeMaker).
2. Ensure you have Go installed on your system.
3. Run `go mod tidy` to install all required dependencies.
4. Place the `.env` file with appropriate values in the root directory or create one based on `.env.example`.

## Usage

```sh
go run main.go -projectName=<projectName> -model=<model> -license -contribute
```

### Parameters:

- `-projectName`: The name of the repository/project for generating the README.
- `-model`: The AI model to use (default is "mistral").
- `-license`: Boolean flag to generate a `LICENSE.txt` file.
- `-contribute`: Boolean flag to generate a `Contribute.md` file.

## Technologies

- **Go**: The programming language used to build the tool.
- **godotenv**: For loading environment variables from the `.env` file.
- **AI Services**: Can utilize different AI models as specified.

## Configuration and Environment

We use a `.env` file for configuring environment variables.

Example `.env` file content:

```bash
API_KEY=""
```

## Folder Structure

```bash
📁 .
├── .env                             # Environment variables
├── .env.example                    # Template for environment variables
├── .gitignore                      # Files and directories to ignore in git
├── Contribute.md                   # Contribution guidelines
├── License.txt                     # License file
├── Readme.md                       # Project documentation
├── cmd                             # Command-line related code
│   ├── contribute.go
│   ├── license.go
│   ├── readme.go
├── go.mod
├── go.sum                         # Go module dependencies
├── internal                        # Internal modules
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
├── main.go                         # Main entry point of the application
```

## Authors

- **WalonCode** (GitHub: [walonCode](https://github.com/walonCode))

## Contribution

Please refer to [Contribute.md](Contribute.md) for guidelines on how to contribute to this project.