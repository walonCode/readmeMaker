package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type ChatCompletionResponse struct {
	Choices []struct {
		Message ChatMessage `json:"message"`
	} `json:"choices"`
}

var models = map[string]string{
	"mistral":"mistral-saba-24b",
	"llama":"llama3-70b-8192",
	"gema":"gemma2-9b-it",
}

func AiGeneration(prompt, model string) (string, error) {
	apiKey := os.Getenv("API_KEY")


	reqBody := ChatCompletionRequest{
		Model: models[model],
		Messages: []ChatMessage{
			{Role: "user", Content: prompt},
		},
	}

	

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST","https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", nil
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res,err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// Handle error responses
	if res.StatusCode != http.StatusOK {
		errorBody, _ := io.ReadAll(res.Body)
		return "", fmt.Errorf("groq API error: %s", string(errorBody))
	}

	var chatResp ChatCompletionResponse
	if err := json.NewDecoder(res.Body).Decode(&chatResp); err != nil {
		return "", err
	}

	if len(chatResp.Choices) == 0 {
		return "", fmt.Errorf("no response from ai")
	}
	return chatResp.Choices[0].Message.Content, nil

}