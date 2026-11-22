package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScanAllFilesExist(t *testing.T) {
	root := t.TempDir()

	files := []string{
		"README.md",
		"LICENSE",
		"go.mod",
	}

	for _, file := range files {
		path := filepath.Join(root, file)

		if err := os.WriteFile(path, []byte("test"), 0644); err != nil {
			t.Fatal(err)
		}
	}

	missing := Scan(root, files)

	if len(missing) != 0 {
		t.Errorf("Expected no missing files, got %v", missing)
	}
}

func TestScanMissingFile(t *testing.T) {
	root := t.TempDir()

	files := []string{
		"README.md",
		"LICENSE",
	}

	if err := os.WriteFile(
		filepath.Join(root, "README.md"),
		[]byte("test"),
		0644,
	); err != nil {
		t.Fatal(err)
	}

	missing := Scan(root, files)

	if len(missing) != 1 {
		t.Fatalf("Expected 1 missing file, got %d", len(missing))
	}

	if missing[0] != "LICENSE" {
		t.Errorf("Expected LICENSE, got %s", missing[0])
	}
}
