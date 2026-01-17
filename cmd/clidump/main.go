package main

import (
	"fmt"
	"os"
	"github.com/mohamedirfansh/clidump/internal/history"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	cmds, err := history.LastN(20)
	if err != nil {
		return err
	}

	for i, cmd := range cmds {
		fmt.Printf("%2d  %s\n", i+1, cmd)
	}
	return nil
}
