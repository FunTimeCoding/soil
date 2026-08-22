package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func findFunction(
	p *packages.Package,
	symbol string,
) (types.Object, *packages.Package, error) {
	o := p.Types.Scope().Lookup(symbol)

	if o == nil {
		return nil, nil, not_found.Format(
			"symbol %s not found in %s",
			symbol,
			p.PkgPath,
		)
	}

	return o, p, nil
}
