package scanner

import (
	"os"
	"path/filepath"
)

func Scan(root string, requiredFiles []string) []string {
	var missing []string

	for _, file := range requiredFiles {
		path := filepath.Join(root, file)

		info, err := os.Stat(path)

		if err != nil || info.IsDir() {
			missing = append(missing, file)
		}
	}

	return missing
}
