package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"path/filepath"
	"testing"
)

func TestNestedModule(t *testing.T) {
	root := t.TempDir()
	system_tester.Write(t, root, "go.mod", "module alfa\n")
	system_tester.Write(t, root, "graft/bravo/go.mod", "module bravo\n")
	system_tester.Write(t, root, "pkg/alfa/alfa.go", "package alfa\n")
	nested := filepath.Join(root, "graft", "bravo")
	assert.Boolean(t, false, system.NestedModule(root, root))
	assert.True(t, system.NestedModule(root, nested))
	assert.Boolean(
		t,
		false,
		system.NestedModule(root, filepath.Join(root, "pkg", "alfa")),
	)
	assert.Boolean(t, false, system.NestedModule(nested, nested))
}
