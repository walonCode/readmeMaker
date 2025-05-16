# ReadmeMaker

ReadmeMaker is a tool designed to automatically generate a comprehensive README.md file for a given project. This tool uses various AI models to create a well-structured and informative README based on the project's directory structure and code snippets.

## Features

- Scans the project directory to detect files and directories.
- Generates a detailed directory structure.
- Uses AI models (Llama, Mistral, Qwen) to generate a README.md based on the project's structure and code snippets.
- Configurable with an environmental variable for the API key.

## Installation

Follow these steps to install and set up the project:

1. **Clone the repository:**
   ```sh
   git clone https://github.com/walonCode/readmeMaker.git
   cd readmeMaker
   ```

2. **Set up the environment:**
   - Copy the example environment file:
     ```sh
     cp .env.example .env
     ```
   - Set your API key in the `.env` file:
     ```sh
     API_KEY="your_api_key_here"
     ```

3. **Install dependencies:**
   ```sh
   go mod tidy
   ```

## Usage

To generate a README.md for your project, follow these steps:

1. **Navigate to your project directory:**
   ```sh
   cd your_project_directory
   ```

2. **Run the ReadmeMaker tool:**
   ```sh
   go run main.go --projectName "Your Project Name" --model "llama"
   ```
   Replace `"Your Project Name"` with the name of your project and `"llama"` with the desired AI model (e.g., "mistral", "qwen").

## Technologies

- **Go:** The main programming language used.
- **godotenv:** A Go package to load environment variables from a `.env` file.

## Configuration and Environment Variables

- `.env` file:
  ```plaintext
  API_KEY=""
  ```
  Set your API key in the `API_KEY` variable in the `.env` file.

- `.env.example`:
  Provided as a template for setting up the environment variables.

## Folder Structure

Here is the organized folder structure of the ReadmeMaker project:

```
readmeMaker/
├── .env
├── .env.example
├── .git/
│   ├── COMMIT_EDITMSG
│   ├── HEAD
│   ├── config
│   ├── description
│   ├── hooks/
│   │   ├── ...
│   ├── index
│   ├── info/
│   │   ├── exclude
│   ├── logs/
│   │   ├── HEAD
│   │   ├── refs/
│   │   ├── remotes/
│   │   ├── tags
│   ├── objects/
│   │   ├──  ...
│   ├── pack
├── .gitignore
├── README.md
├── go.mod
├── go.sum
├── internal/
│   ├── file_parser/
│   │   ├── fileParser.go
│   ├── llm/
│   │   ├── ai.services.go
│   ├── tree/
│   │   ├── buildTree.go
│   ├── utils/
│   │   ├── buildPrompt.go
│   │   ├── writeReadme.go
├── main.go
```

## Authors

- **Mohamed Lamin Walon-Jalloh** (https://www.linkedin.com/in/mohamed-walon/)

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes and commit them (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Create a Pull Request.