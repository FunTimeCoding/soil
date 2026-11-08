package face

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"github.com/funtimecoding/soil/pkg/source/index"
	"golang.org/x/tools/go/packages"
)

func FromWorkspace(
	w *index.Workspace,
	loaded []*packages.Package,
) *Set {
	result := NewWithIndex(w, loaded)
	live := make(map[string]bool, len(loaded))

	for _, p := range loaded {
		live[p.PkgPath] = true
	}

	for path, interfaces := range index.Facts[*[]*fact.Interface](w, Kind()) {
		if !live[path] {
			result.Add(*interfaces)
		}
	}

	return result
}
