# ReadmeMaker

## Description

ReadmeMaker is a project designed to automatically generate custom readable markdown files based on specific attributes and codebase structures. It leverages several internal modules to facilitate this process, enhancing productivity and ensuring consistency across documentation.

## Features

- Automatically generates `README.md` files
- Parses project files to extract relevant information
- Uses AI services to tailor content
- Integrated environment variable management

## Installation

To install ReadmeGenerator, follow these steps:

1. Clone the repository:
   ```sh
   git clone https://github.com/walonCode/readmeMaker.git
   ```

2. Navigate into the project directory:
   ```sh
   cd readmeMaker
   ```

3. Install dependencies:
   ```sh
   go mod tidy
   ```

## Usage

To use ReadmeGenerator, follow these instructions:

1. Configure your environment variables in the `.env` file.
2. Run the main executable:
   ```sh
   go run main.go
   ```

## Technologies

- **Programming Language**: Go
- **Dependencies**: joho/godotenv

## Configuration and Environment Variables

Create a `.env` file in the root directory, and add your API key:

```
API_KEY=your_api_key_here
```

## Folder Structure

```markdown
.
├── .env
├── .env.example
├── .git
│   ├── ...
├── .gitignore
├── Readme.md
├── go.mod
├── go.sum
├── internal
│   ├── file_parser
│   │   ├── fileParser.go
│   ├── llm
│   │   ├── ai.services.go
│   ├── tree
│   │   ├── buildTree.go
│   ├── utils
│   │   ├── buildPrompt.go
│   │   ├── writeReadme.go
├── main.go
```

## Authors

- **John Doe** ([@n4tur4l](https://github.com/n4tur4l))

## Contribution

Contributions to ReadmeGenerator are welcome. To contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch").
3. Make your changes and push them (`git push origin feature-branch`).
4. Open a Pull Request.