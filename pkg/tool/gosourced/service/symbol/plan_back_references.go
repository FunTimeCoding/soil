package symbol

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/identifier"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/relocation"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"slices"
	"sort"
)

func planBackReferences(
	all []*packages.Package,
	p *packages.Package,
	entries []*relocation.Entry,
	skip map[*ast.Ident]bool,
) ([]types.Object, map[*ast.Ident]string, error) {
	excluded := make(map[token.Pos]bool)

	for _, entry := range entries {
		excluded[entry.Object.Pos()] = true
	}

	seen := make(map[types.Object]bool)
	var flipped []types.Object

	for _, entry := range entries {
		entry.BackIdentifiers = slices.DeleteFunc(
			backReferenceIdentifiers(p, excluded, entry.Node),
			func(i *ast.Ident) bool { return skip[i] },
		)

		for _, ident := range entry.BackIdentifiers {
			o := p.TypesInfo.Uses[ident]

			if o.Exported() || seen[o] {
				continue
			}

			seen[o] = true
			flipped = append(flipped, o)
		}
	}

	sort.Slice(
		flipped,
		func(
			i int,
			j int,
		) bool {
			return flipped[i].Pos() < flipped[j].Pos()
		},
	)
	renames := make(map[*ast.Ident]string)

	for _, o := range flipped {
		name := identifier.FlipName(o.Name())

		if f := identifier.CheckScopeCollision(p, name); f != nil {
			return nil, nil, f
		}

		for ident, d := range p.TypesInfo.Defs {
			if d == o {
				renames[ident] = name
			}
		}

		for _, f := range resolve.FindAllReferences(all, o) {
			renames[f.Ident] = name
		}

		if capture := embeddedRenames(all, o, name, renames); capture != "" {
			return nil, nil, conflict.Format("%s", capture)
		}
	}

	return flipped, renames, nil
}
