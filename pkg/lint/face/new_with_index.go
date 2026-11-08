package face

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"github.com/funtimecoding/soil/pkg/source/index"
	"golang.org/x/tools/go/packages"
)

func NewWithIndex(
	w *index.Workspace,
	loaded []*packages.Package,
) *Set {
	result := New(loaded)

	for _, interfaces := range index.Externals[*[]*fact.Interface](w, Kind()) {
		result.Add(*interfaces)
	}

	return result
}
