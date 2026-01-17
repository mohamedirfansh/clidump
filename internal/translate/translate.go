package translate

import (
    "context"
    "fmt"
    "os"

    "github.com/sashabaranov/go-openai"
)

// ToCommand translates natural language to Unix commands using OpenAI
func ToCommand(naturalLanguage string) (string, error) {
    apiKey := os.Getenv("CLIDUMP_OPENAI_KEY")
    if apiKey == "" {
        return "", fmt.Errorf("CLIDUMP_OPENAI_KEY environment variable not set")
    }

    client := openai.NewClient(apiKey)
    
    resp, err := client.CreateChatCompletion(
        context.Background(),
        openai.ChatCompletionRequest{
            Model: openai.GPT4o,
            Messages: []openai.ChatCompletionMessage{
                {
                    Role:    "system",
                    Content: "You are a Unix command translator. Convert natural language descriptions to Unix commands. Return ONLY the command, no explanations or markdown.",
                },
                {
                    Role:    "user",
                    Content: naturalLanguage,
                },
            },
            Temperature: 0.3,
            MaxTokens:   150,
        },
    )

    if err != nil {
        return "", fmt.Errorf("OpenAI API error: %w", err)
    }

    if len(resp.Choices) == 0 {
        return "", fmt.Errorf("no response from OpenAI")
    }

    return resp.Choices[0].Message.Content, nil
}