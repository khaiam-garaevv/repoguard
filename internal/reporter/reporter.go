package reporter

import "fmt"

func Report(total int, missing []string) {
	passed := total - len(missing)

	fmt.Println("--------------------------------")
	fmt.Printf("Total:  %d\n", total)
	fmt.Printf("Passed: %d\n", passed)
	fmt.Printf("Failed: %d\n", len(missing))

	if len(missing) > 0 {
		fmt.Println("\nMissing files:")

		for _, file := range missing {
			fmt.Println("[FAIL]", file)
		}
	}
}
