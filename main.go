package main

import (
	"fmt"
	"os"

	"github.com/khayyamgarayev/repoguard/internal/checker"
	"github.com/khayyamgarayev/repoguard/internal/reporter"
	"github.com/khayyamgarayev/repoguard/internal/scanner"
)

func main() {
	root := "."

	if len(os.Args) > 1 {
		root = os.Args[1]
	}

	required := checker.RequiredFiles()

	fmt.Println("RepoGuard - Repository Analyzer")
	fmt.Println("--------------------------------")

	missing := scanner.Scan(root, required)

	reporter.Report(len(required), missing)

	if len(missing) > 0 {
		os.Exit(1)
	}
}
