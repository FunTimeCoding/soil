package kind_tester

import (
	"github.com/funtimecoding/soil/pkg/source/index/kind"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func Listed(name string) *kind.Kind {
	return kind.New(
		name,
		func() any {
			return new([]string)
		},
		func(p *packages.Package) any {
			var none []string

			return &none
		},
		func(p *types.Package) any {
			return new([]string{p.Path()})
		},
	)
}
