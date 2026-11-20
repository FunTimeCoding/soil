package module

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"testing"
)

func New(
	t *testing.T,
	name string,
) *Module {
	t.Helper()
	directory := t.TempDir()
	system_tester.New(t).Write(
		directory,
		"go.mod",
		fmt.Sprintf("module %s\n\ngo 1.22\n", name),
	)

	return &Module{t: t, directory: directory}
}
