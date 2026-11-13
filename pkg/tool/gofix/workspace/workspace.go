package workspace

import (
	"github.com/funtimecoding/soil/pkg/source/layer"
	"go/token"
	"golang.org/x/tools/go/packages"
)

type Workspace struct {
	diff         bool
	stacked      bool
	roots        []string
	layer        *layer.Layer
	loadedKey    string
	loadedWrites int
	loaded       []*packages.Package
	loadedSet    *token.FileSet
}
