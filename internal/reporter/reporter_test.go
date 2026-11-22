package reporter

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestReport(t *testing.T) {
	missing := []string{"LICENSE"}

	oldStdout := os.Stdout

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	os.Stdout = w

	Report(4, missing)

	w.Close()
	os.Stdout = oldStdout

	var output bytes.Buffer

	if _, err := io.Copy(&output, r); err != nil {
		t.Fatal(err)
	}

	result := output.String()

	expected := []string{
		"Total:  4",
		"Passed: 3",
		"Failed: 1",
		"[FAIL] LICENSE",
	}

	for _, text := range expected {
		if !strings.Contains(result, text) {
			t.Errorf("Expected output to contain %q, got:\n%s", text, result)
		}
	}
}
