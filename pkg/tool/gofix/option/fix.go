package option

import (
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/layer"
	"go/token"
	"golang.org/x/tools/go/packages"
)

type Fix struct {
	Root         string
	Patterns     []string
	Diff         bool
	Full         bool
	Ours         []string
	Index        string
	Replacing    []string
	Reported     map[string]bool
	Handle       *index.Workspace
	Loaded       []*packages.Package
	FileSet      *token.FileSet
	Layer        *layer.Layer
	BeforeCommit func()
}
