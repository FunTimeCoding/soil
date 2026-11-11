package identifier

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/source/types/resolve_reference"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func EmbeddedFieldReferences(
	all []*packages.Package,
	o types.Object,
	newName string,
) ([]resolve_reference.Reference, string) {
	typeName, okay := o.(*types.TypeName)

	if !okay {
		return nil, ""
	}

	var result []resolve_reference.Reference
	seen := make(map[token.Pos]bool)

	for _, p := range all {
		for _, d := range p.TypesInfo.Defs {
			v, okay := d.(*types.Var)

			if !okay || !v.Embedded() || v.Name() != typeName.Name() ||
				seen[v.Pos()] || !embedsDirectly(v.Type(), typeName) {
				continue
			}

			seen[v.Pos()] = true

			if holder := embeddingCapture(p, v, newName); holder != "" {
				return nil, fmt.Sprintf(
					"%s already reaches %s - renaming its embedded %s would capture it",
					holder,
					newName,
					typeName.Name(),
				)
			}

			for _, f := range resolve.FindAllReferences(all, v) {
				if f.Package.TypesInfo.Defs[f.Ident] == nil {
					result = append(result, f)
				}
			}
		}
	}

	return result, ""
}
