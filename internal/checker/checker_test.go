package checker

import (
	"reflect"
	"testing"
)

func TestRequiredFiles(t *testing.T) {
	expected := []string{
		"README.md",
		".gitignore",
		"go.mod",
		"LICENSE",
	}

	actual := RequiredFiles()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(
			"Expected %v, got %v",
			expected,
			actual,
		)
	}
}
