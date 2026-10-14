package kind_tester

import (
	"github.com/funtimecoding/soil/pkg/source/index/kind"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"golang.org/x/tools/go/packages"
)

func CountingSyntax(
	name string,
	external *int,
) *kind.Kind {
	return kind.New(
		name,
		func() any {
			return new(string)
		},
		func(p *packages.Package) any {
			return new(p.PkgPath)
		},
		nil,
	).WithSyntaxExternal(
		func(p *packages.Package) any {
			*external++

			if p.TypesInfo == nil || len(p.Syntax) == 0 {
				return new(join.Empty("!", p.PkgPath))
			}

			return new(p.PkgPath)
		},
	)
}
