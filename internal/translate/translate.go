package translate

import (
    "context"
    "fmt"
    "os"

    "github.com/sashabaranov/go-openai"
)

const (
    groqBaseURL = "https://api.groq.com/openai/v1"
    groqModel   = "llama-3.1-8b-instant"
)

// ToCommand translates natural language to Unix commands using Groq API
func ToCommand(naturalLanguage string) (string, error) {
    apiKey := os.Getenv("CLIDUMP_GROQ_KEY")
    if apiKey == "" {
        return "", fmt.Errorf("CLIDUMP_GROQ_KEY environment variable not set")
    }

    // Create Groq client using OpenAI-compatible SDK
    config := openai.DefaultConfig(apiKey)
    config.BaseURL = groqBaseURL
    client := openai.NewClientWithConfig(config)
    
    resp, err := client.CreateChatCompletion(
        context.Background(),
        openai.ChatCompletionRequest{
            Model: groqModel,
            Messages: []openai.ChatCompletionMessage{
                {
                    Role:    openai.ChatMessageRoleSystem,
                    Content: "You are a Unix command translator. Convert natural language descriptions to Unix commands. Return ONLY the command, no explanations or markdown.",
                },
                {
                    Role:    openai.ChatMessageRoleUser,
                    Content: naturalLanguage,
                },
            },
            Temperature: 0.3,
            MaxTokens:   150,
        },
    )

    if err != nil {
        return "", fmt.Errorf("Groq API error: %w", err)
    }

    if len(resp.Choices) == 0 {
        return "", fmt.Errorf("no response from Groq")
    }

    return resp.Choices[0].Message.Content, nil
}