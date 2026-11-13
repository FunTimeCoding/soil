package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/layer"
	"path/filepath"
	"testing"
)

func TestLayerChangesSpanStack(t *testing.T) {
	root := layerModule(t)
	base := layer.New(nil, root)
	alfa := filepath.Join(root, "alfa/alfa.go")
	base.Write(alfa, []byte("package alfa\n\n"))
	child := layer.New(base)
	foxtrot := filepath.Join(root, "alfa/foxtrot.go")
	child.Move(alfa, foxtrot)
	charlie := filepath.Join(root, "charlie/charlie.go")
	child.Write(charlie, []byte("package charlie\n\n"))
	changes := child.Changes()
	assert.Count(t, 2, changes)
	assert.String(t, alfa, changes[0].Path)
	assert.String(t, foxtrot, changes[0].Destination)
	assert.String(t, "package alfa\n", string(changes[0].Original))
	assert.String(t, "package alfa\n\n", string(changes[0].Current))
	assert.String(t, charlie, changes[1].Path)
}
