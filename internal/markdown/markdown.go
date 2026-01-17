package markdown

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Generate creates a markdown file with commands and their explanations
func Generate(commands []string, explanations map[string]string, outputDir string) (string, error) {
	if outputDir == "" {
		// Default to current directory
		var err error
		outputDir, err = os.Getwd()
		if err != nil {
			return "", fmt.Errorf("failed to get working directory: %w", err)
		}
	}

	// Find the next available filename
	filename, err := getNextFilename(outputDir)
	if err != nil {
		return "", fmt.Errorf("failed to generate filename: %w", err)
	}

	// Build markdown content
	var content strings.Builder
	content.WriteString("# CLI Command History\n\n")

	for _, cmd := range commands {
		explanation := explanations[cmd]
		if explanation == "" {
			explanation = "No explanation available"
		}

		content.WriteString(fmt.Sprintf("## `%s`\n\n", cmd))
		content.WriteString(fmt.Sprintf("%s\n\n", explanation))
	}

	// Write to file
	filepath := filepath.Join(outputDir, filename)
	err = os.WriteFile(filepath, []byte(content.String()), 0644)
	if err != nil {
		return "", fmt.Errorf("failed to write markdown file: %w", err)
	}

	return filepath, nil
}

// getNextFilename finds the next available filename (clidump-1.md, clidump-2.md, etc.)
func getNextFilename(dir string) (string, error) {
	for i := 1; ; i++ {
		filename := fmt.Sprintf("clidump-%d.md", i)
		filepath := filepath.Join(dir, filename)

		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			return filename, nil
		}
	}
}
