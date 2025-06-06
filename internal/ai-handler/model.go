package aihandler

import (
	"os"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

const OpenAIEndpoint = "https://api.openai.com/v1/chat/completions"

var OpenAIAPIKey = os.Getenv("OPENAI_API_KEY")
