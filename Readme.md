# ReadmeMaker

### Description

ReadmeMaker is a command-line tool for generating professional and clean README files, as well as contribution guidelines and licenses, for your projects. By integrating artificial intelligence, ReadmeMaker simplifies the process of maintaining up-to-date documentation, saving you valuable time and effort.

### Features

- Generate README files tailored for different types of projects.
- Include licenses for open-source projects.
- Generate contribution guidelines.
- Support for multiple AI models for generating content.

### Installation

1. **Clone the repository:**

    ```bash
    git clone github.com/walonCode/readmeMaker
    cd readmeMaker
    ```

2. **Install dependencies:**

    ```bash
    go get
    ```

3. **Set up environment variables:**

    Copy `.env.example` to `.env` and fill in the required values.

    ```bash
    cp .env.example .env
    ```

    Populate `.env` with your API key and other necessary variables.

### Usage

Run the tool using the following command:

```bash
go run main.go -projectName <projectName> -model <model> -license <license type> -contribute
```

- `projectName`: The name of the project/repository.
- `model`: The AI model to use for content generation.
- `license`: The type of license you want to generate.
- `contribute`: Option to include contribution guidelines (boolean flag).

### Technologies

| Technology/Tool                     | Purpose                                                                                           |
|-------------------------------------|---------------------------------------------------------------------------------------------------|
| ![Go](https://github.githubassets.com/images/icons/emoji/go.png) Go | Programming Language                                                         |
| ![godotenv](https://github.githubassets.com/images/icons/emoji/env.png) godotenv | Loading environment variables                                                              |
| ![AI Services](https://github.githubassets.com/images/icons/emoji/robot.png) AI Services | Generating project documentation automatically using AI                                      |
| ![File Parsing](https://github.githubassets.com/images/icons/emoji/file.png) File Parsing | Parsing and writing project-related files                                                  |

### Configuration and Environment Variables

- **.env:**

    ```plaintext
    API_KEY=""
    ```

### Folder Structure

```
.
├── .env
├── .env.example
├── .gitignore
├── cmd
|  ├── contribute.go
|  ├── license.go
|  ├── readme.go
├── go.mod
├── go.sum
├── internal
|  ├── file_parser
|  |  ├── fileParser.go
|  ├── llm
|  |  ├── ai.services.go
|  ├── tree
|  |  ├── buildTree.go
|  ├── types
|  |  ├── types.go
|  ├── utils
|  |  ├── buildPrompt.go
|  |  ├── writeReadme.go
├── main.go
```

### Author

- **walonCode**

---

Feel free to raise an issue or submit a PR if you'd like to contribute to the project.

---