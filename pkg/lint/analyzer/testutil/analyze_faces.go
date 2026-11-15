package testutil

import (
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"golang.org/x/tools/go/packages"
	"testing"
)

func AnalyzeFaces(
	t *testing.T,
	directory string,
	check func(
		*packages.Package,
		*output.Results,
		*face.Set,
	),
) *output.Results {
	t.Helper()
	loaded, result := loadVariants(t, directory)
	faces := face.New(loaded)

	for _, p := range loaded {
		check(p, result, faces)
	}

	return result
}
