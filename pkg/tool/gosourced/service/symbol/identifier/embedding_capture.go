package identifier

import (
	"go/types"
	"golang.org/x/tools/go/packages"
)

func embeddingCapture(
	p *packages.Package,
	field *types.Var,
	newName string,
) string {
	for _, o := range p.TypesInfo.Defs {
		holder, okay := o.(*types.TypeName)

		if !okay || holder.IsAlias() {
			continue
		}

		s, okay := holder.Type().Underlying().(*types.Struct)

		if !okay || !structHolds(s, field) {
			continue
		}

		if found, _, _ := types.LookupFieldOrMethod(
			types.NewPointer(holder.Type()),
			true,
			holder.Pkg(),
			newName,
		); found != nil {
			return holder.Name()
		}
	}

	return ""
}
