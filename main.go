package main

import (
	"fmt"
	"os"

	"github.com/khaiam-garaevv/repoguard/internal/checker"
	"github.com/khaiam-garaevv/repoguard/internal/reporter"
	"github.com/khaiam-garaevv/repoguard/internal/scanner"
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
