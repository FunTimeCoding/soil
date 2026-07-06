package testutil

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"golang.org/x/tools/go/packages"
	"testing"
)

func LoadTestPackage(
	t *testing.T,
	sourceDirectory string,
	extraModLines ...string,
) (*packages.Package, *output.Results) {
	t.Helper()
	temporary := PrepareTestPackage(t, sourceDirectory, extraModLines...)

	return LoadFromDirectory(t, temporary)
}
