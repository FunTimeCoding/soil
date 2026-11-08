package kind_tester

import (
	"github.com/funtimecoding/soil/pkg/source/index/kind"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/types"
	"golang.org/x/tools/go/packages"
	"slices"
)

func ScopeNames(name string) *kind.Kind {
	return kind.New(
		name,
		func() any {
			return new(string)
		},
		func(p *packages.Package) any {
			return new(p.PkgPath)
		},
		func(p *types.Package) any {
			return new(
				join.Space(
					slices.Concat([]string{p.Path()}, p.Scope().Names())...),
			)
		},
	)
}
