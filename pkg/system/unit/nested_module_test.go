package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
	"testing"
)

func TestNestedModule(t *testing.T) {
	root := t.TempDir()
	testutil.WriteFile(t, root, "go.mod", "module alfa\n")
	testutil.WriteFile(t, root, "graft/bravo/go.mod", "module bravo\n")
	testutil.WriteFile(t, root, "pkg/alfa/alfa.go", "package alfa\n")
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
