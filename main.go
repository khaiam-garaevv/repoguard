package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	target := "."

	if len(os.Args) > 1 {
		target = os.Args[1]
	}

	requiredFiles := []string{
		"README.md",
		".gitignore",
		"go.mod",
	}

	fmt.Println("RepoGuard - Repository Analyzer")
	fmt.Println("--------------------------------")

	missing := 0

	for _, file := range requiredFiles {
		path := filepath.Join(target, file)

		info, err := os.Stat(path)

		if err == nil && !info.IsDir() {
			fmt.Println("[OK] ", file)
		} else if os.IsNotExist(err) {
			fmt.Println("[FAIL]", file)
			missing++
		} else {
			fmt.Println("[ERROR]", file)
			missing++
		}
	}

	fmt.Println("--------------------------------")
	fmt.Printf("Total: %d\n", len(requiredFiles))
	fmt.Printf("Passed: %d\n", len(requiredFiles)-missing)
	fmt.Printf("Failed: %d\n", missing)

	if missing > 0 {
		os.Exit(1)
	}
}
