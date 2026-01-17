package openai

import (
	"context"
	"fmt"
	"strings"

	"github.com/sashabaranov/go-openai"
)

// ExplainCommands takes a list of CLI commands and returns their explanations using OpenAI
func ExplainCommands(apiKey string, commands []string) (map[string]string, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("OpenAI API key is required")
	}

	client := openai.NewClient(apiKey)
	explanations := make(map[string]string)

	// Build a single prompt with all commands for efficiency
	var promptBuilder strings.Builder
	promptBuilder.WriteString("Provide a succinct one-sentence explanation for each of the following CLI commands.\n")
	promptBuilder.WriteString("Format your response as a numbered list matching the input, with just the explanation for each command.\n\n")

	for i, cmd := range commands {
		promptBuilder.WriteString(fmt.Sprintf("%d. %s\n", i+1, cmd))
	}

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: promptBuilder.String(),
				},
			},
			MaxTokens:   500,      // ↓ also reduce this
			Temperature: 0.3,
		},
	)

	if err != nil {
		return nil, fmt.Errorf("OpenAI API error: %w", err)
	}

	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("no response from OpenAI")
	}

	// Parse the numbered response
	response := resp.Choices[0].Message.Content
	lines := strings.Split(response, "\n")

	cmdIndex := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Try to match numbered format: "1. explanation" or "1) explanation"
		// Handle multi-digit numbers (10., 11., etc.)
		dotIdx := strings.Index(line, ".")
		parenIdx := strings.Index(line, ")")

		separatorIdx := -1
		if dotIdx > 0 && dotIdx < 4 { // Reasonable number length
			separatorIdx = dotIdx
		} else if parenIdx > 0 && parenIdx < 4 {
			separatorIdx = parenIdx
		}

		if separatorIdx > 0 && cmdIndex < len(commands) {
			// Extract explanation after the separator
			explanation := strings.TrimSpace(line[separatorIdx+1:])
			explanations[commands[cmdIndex]] = explanation
			cmdIndex++
		}
	}

	// Ensure all commands have explanations
	for _, cmd := range commands {
		if _, exists := explanations[cmd]; !exists {
			explanations[cmd] = "Command explanation not available"
		}
	}

	return explanations, nil
}
