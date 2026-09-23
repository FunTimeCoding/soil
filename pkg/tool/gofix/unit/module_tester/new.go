package module_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func New(
	t *testing.T,
	name string,
) *Module {
	t.Helper()
	directory := t.TempDir()
	testutil.WriteFile(
		t,
		directory,
		"go.mod",
		fmt.Sprintf("module %s\n\ngo 1.22\n", name),
	)

	return &Module{t: t, directory: directory}
}
