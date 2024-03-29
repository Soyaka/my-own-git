package commands

import (
	"fmt"
	"os"
)

func GitInit() error {
	for _, dir := range []string{"git", "git/objects", "git/refs"} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating directory: %s\n", err)
			return err
		}
	}
	headFileContents := []byte("ref: refs/heads/master\n")
	if err := os.WriteFile("git/HEAD", headFileContents, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %s\n", err)
		return err
	}

	fmt.Println("Initialized git directory")
	return nil
}
