package unit

import (
	"github.com/funtimecoding/soil/pkg/source/index/kind"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func countingKind(
	name string,
	extracted *int,
	external *int,
) *kind.Kind {
	return kind.New(
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
