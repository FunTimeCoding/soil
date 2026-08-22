package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func findMethod(
	p *packages.Package,
	symbol string,
	receiver string,
) (types.Object, *packages.Package, error) {
	o := p.Types.Scope().Lookup(receiver)

	if o == nil {
		return nil, nil, not_found.Format(
			"receiver type %s not found in %s",
			receiver,
			p.PkgPath,
		)
	}

	named, okay := o.Type().(*types.Named)

	if !okay {
		return nil, nil, validation.New(
			"%s is not a named type in %s",
			receiver,
			p.PkgPath,
		)
	}

	for i := range named.NumMethods() {
		m := named.Method(i)

		if m.Name() == symbol {
			return m, p, nil
		}
	}

	return nil, nil, not_found.Format(
		"method %s not found on %s in %s",
		symbol,
		receiver,
		p.PkgPath,
	)
}
