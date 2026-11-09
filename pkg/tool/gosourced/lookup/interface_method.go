package lookup

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func interfaceMethod(
	i *types.Interface,
	p *packages.Package,
	symbol string,
	receiver string,
) (types.Object, *packages.Package, error) {
	for m := range i.ExplicitMethods() {
		if m.Name() == symbol {
			return m, p, nil
		}
	}

	for m := range i.Methods() {
		if m.Name() == symbol {
			return nil, nil, validation.New(
				"%s is declared in %s, embedded in %s - rename it there",
				symbol,
				types.TypeString(
					m.Type().(*types.Signature).Recv().Type(),
					nil,
				),
				receiver,
			)
		}
	}

	return nil, nil, not_found.Format(
		"method %s not found on %s in %s",
		symbol,
		receiver,
		p.PkgPath,
	)
}
