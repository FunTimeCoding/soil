package kind_tester

import (
	"github.com/funtimecoding/soil/pkg/source/index/kind"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func Declared(name string) *kind.Kind {
	return kind.New(
		name,
		func() any {
			return new(string)
		},
		func(p *packages.Package) any {
			return new(join.Space(p.Types.Scope().Names()...))
		},
		func(p *types.Package) any {
			return new(p.Path())
		},
	)
}
