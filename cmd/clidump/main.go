package main

import (
	"fmt"
	"os"

	"github.com/mohamedirfansh/clidump/internal/history"
)

const (
	DEFAULT_COMMANDS_TO_DUMP = 20
)

func main() {
	/**
	* If none of the commands match, the default is to dump the last
	* DEFAULT_COMMANDS_TO_DUMP number of commands on terminal
	 */
	if err := dumpLatestCommands(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func dumpLatestCommands() error {
	cmds, err := history.LastN(DEFAULT_COMMANDS_TO_DUMP)
	if err != nil {
		return err
	}

	for i, cmd := range cmds {
		fmt.Printf("%2d  %s\n", i+1, cmd)
	}
	return nil
}
