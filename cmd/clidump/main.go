package main

import (
	"flag"
	"fmt"
	"os"
    "path/filepath"
    "strings"

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
	install := flag.Bool("install", false, "Install shell integration")
    flag.Parse()

	// Handle --install flag
    if *install {
        if err := installShellIntegration(); err != nil {
            fmt.Fprintf(os.Stderr, "Error: %v\n", err)
            os.Exit(1)
        }
        return
    }

    // If -t flag is provided, translate and exit
    if *englishCmd != "" {
        if err := translateCommand(*englishCmd); err != nil {
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

func installShellIntegration() error {
    // Detect shell
    shell := os.Getenv("SHELL")
    var shellRC string
    var wrapperFunc string

    if strings.Contains(shell, "zsh") {
        shellRC = filepath.Join(os.Getenv("HOME"), ".zshrc")
        wrapperFunc = `
# clidump shell integration
ct() {
    if [ -z "$1" ]; then
        echo "Usage: ct \"description of command\""
        return 1
    fi
    
    local cmd=$(clidump -t "$*")
    
    if [ $? -eq 0 ] && [ -n "$cmd" ]; then
        print -z "$cmd"
    fi
}
`
   } else if strings.Contains(shell, "bash") {
        shellRC = filepath.Join(os.Getenv("HOME"), ".bashrc")
        wrapperFunc = `
# clidump shell integration
ct() {
    if [ -z "$1" ]; then
        echo "Usage: ct \"description of command\""
        return 1
    fi
    
    local cmd=$(clidump -t "$*")
    
    if [ $? -eq 0 ] && [ -n "$cmd" ]; then
        bind '"\e[0n": "'"$cmd"'"'
        bind '"\e[0n"'
    fi
}
`
    } else {
        return fmt.Errorf("unsupported shell: %s", shell)
    }

    // Check if already installed
    content, err := os.ReadFile(shellRC)
    if err != nil && !os.IsNotExist(err) {
        return fmt.Errorf("failed to read %s: %w", shellRC, err)
    }

    if strings.Contains(string(content), "# clidump shell integration") {
        fmt.Printf("✓ Shell integration already installed in %s\n", shellRC)
        fmt.Printf("\nRun: source %s\n", shellRC)
        return nil
    }

    // Append wrapper function
    f, err := os.OpenFile(shellRC, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return fmt.Errorf("failed to open %s: %w", shellRC, err)
    }
    defer f.Close()

    if _, err := f.WriteString(wrapperFunc); err != nil {
        return fmt.Errorf("failed to write to %s: %w", shellRC, err)
    }

    fmt.Printf("✓ Shell integration installed in %s\n", shellRC)
    fmt.Printf("\nTo activate, run:\n")
    fmt.Printf("  source %s\n\n", shellRC)
    fmt.Printf("Then use:\n")
    fmt.Printf("  ct \"list all files sorted by size\"\n")
    
    return nil
}

func translateCommand(englishDesc string) error {
    command, err := translate.ToCommand(englishDesc)
    if err != nil {
        return err
    }

	fmt.Printf("%s", command)    
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
