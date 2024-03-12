package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/git-starter-go/cmd/mygit/commands"
)

// Usage: your_git.sh <command> <arg1> <arg2> ...
func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: mygit <command> [<args>...]\n")
		os.Exit(1)
	}
	switch command := os.Args[1]; command {
	case "init":
		err := commands.GitInit()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
	case "cat-file":
		hash := os.Args[2]
		data, err := commands.CatFile(hash)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		fmt.Print(data)

	default:
		fmt.Fprintf(os.Stderr, "Unknown command %s\n", command)
		os.Exit(1)
	}
}
