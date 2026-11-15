package testutil

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"golang.org/x/tools/go/packages"
	"testing"
)

func Analyze(
	t *testing.T,
	directory string,
	check func(*packages.Package, *output.Results),
) *output.Results {
	t.Helper()
	loaded, result := loadVariants(t, directory)

	for _, p := range loaded {
		check(p, result)
	}

	return result
}
