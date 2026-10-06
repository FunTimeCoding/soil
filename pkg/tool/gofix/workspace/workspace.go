package workspace

import (
	"github.com/funtimecoding/soil/pkg/source/snapshot"
	"go/token"
	"golang.org/x/tools/go/packages"
)

type Workspace struct {
	diff         bool
	roots        []string
	before       *snapshot.Snapshot
	beforeCommit func()
	overlay      map[string][]byte
	original     map[string][]byte
	order        []string
	writes       int
	loadedKey    string
	loadedWrites int
	loaded       []*packages.Package
	loadedSet    *token.FileSet
}
