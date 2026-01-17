package history

import (
	"bufio"
	"os"
	"path/filepath"
)

func LastN(n int) ([]string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	files := []string{
		filepath.Join(home, ".zsh_history"),
		filepath.Join(home, ".bash_history"),
	}

	var file *os.File
	for _, f := range files {
		if _, err := os.Stat(f); err == nil {
			file, err = os.Open(f)
			if err != nil {
				return nil, err
			}
			defer file.Close()
			break
		}
	}

	if file == nil {
		return nil, nil
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, parse(scanner.Text()))
	}

	if len(lines) <= n {
		return lines, nil
	}

	return lines[len(lines)-n:], nil
}

// LastNUnique returns the last n unique commands from shell history
func LastNUnique(n int) ([]string, error) {
	// Fetch n commands
	commands, err := LastN(n)
	if err != nil {
		return nil, err
	}

	// Deduplicate while preserving order (keep the most recent occurrence)
	seen := make(map[string]bool)
	var unique []string

	// Iterate in reverse to keep the most recent occurrence
	for i := len(commands) - 1; i >= 0; i-- {
		cmd := commands[i]
		if cmd != "" && !seen[cmd] {
			seen[cmd] = true
			unique = append([]string{cmd}, unique...)
		}
	}

	// // Return the last n unique commands
	// if len(unique) > n {
	// 	return unique[len(unique)-n:], nil
	// }
	return unique, nil
}

// zsh has metadata like ": 123456:0;command"
func parse(line string) string {
	for i, ch := range line {
		if ch == ';' {
			return line[i+1:]
		}
	}
	return line
}
