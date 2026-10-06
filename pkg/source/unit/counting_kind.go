package unit

import (
	"github.com/funtimecoding/soil/pkg/source/index"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func countingKind(
	name string,
	extracted *int,
	external *int,
) *index.Kind {
	return index.NewKind(
		name,
		func() any {
			return new(string)
		},
		func(p *packages.Package) any {
			*extracted++

			return new(p.PkgPath)
		},
		func(p *types.Package) any {
			*external++

			return new(p.Path())
		},
	)
}
