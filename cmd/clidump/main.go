package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mohamedirfansh/clidump/internal/history"
	"github.com/mohamedirfansh/clidump/internal/markdown"
	"github.com/mohamedirfansh/clidump/internal/openai"
	"github.com/mohamedirfansh/clidump/internal/translate"
)

const (
	DEFAULT_COMMANDS_TO_DUMP = 20
)

func main() {
    // Define flags
    englishCmd := flag.String("t", "", "Translate English description to Unix command")
    verbose := flag.Bool("verbose", false, "Show translation explanation")
    flag.Parse()

    // If -t flag is provided, translate and exit
    if *englishCmd != "" {
        if err := translateCommand(*englishCmd, *verbose); err != nil {
            fmt.Fprintf(os.Stderr, "Error: %v\n", err)
            os.Exit(1)
        }
        return
    }

    if err := generateMarkdownDump(); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}

func translateCommand(englishDesc string, verbose bool) error {
    command, err := translate.ToCommand(englishDesc)
    if err != nil {
        return err
    }

	fmt.Printf("\nSuggested command:\n%s\n", command)

    // if verbose {
    //     fmt.Printf("Translating: %s\n", englishDesc)
    //     fmt.Printf("\nSuggested command:\n%s\n", command)
    // } else {
    //     // Just output the command directly (no newline)
    //     fmt.Printf("%s", command)
    // }
    
    return nil
}

func generateMarkdownDump() error {
	// Get Groq API key from environment
	apiKey := os.Getenv("CLIDUMP_GROQ_KEY")
	if apiKey == "" {
		return fmt.Errorf("CLIDUMP_GROQ_KEY environment variable not set")
	}	

	// Get the last 20 unique commands
	fmt.Println("Fetching command history...")
	commands, err := history.LastNUnique(DEFAULT_COMMANDS_TO_DUMP)
	if err != nil {
		return fmt.Errorf("failed to fetch command history: %w", err)
	}

	if len(commands) == 0 {
		return fmt.Errorf("no commands found in history")
	}

	fmt.Printf("Found %d unique commands\n", len(commands))

	// Get explanations from Groq
	fmt.Println("Generating explanations with Groq...")
	explanations, err := openai.ExplainCommands(apiKey, commands)
	if err != nil {
		return fmt.Errorf("failed to get command explanations: %w", err)
	}

	// Generate markdown file
	fmt.Println("Creating markdown file...")
	filepath, err := markdown.Generate(commands, explanations, "")
	if err != nil {
		return fmt.Errorf("failed to generate markdown: %w", err)
	}

	fmt.Printf("✓ Successfully created %s\n", filepath)
	return nil
}
