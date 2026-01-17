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

// zsh has metadata like ": 123456:0;command"
func parse(line string) string {
	for i, ch := range line {
		if ch == ';' {
			return line[i+1:]
		}
	}
	return line
}
