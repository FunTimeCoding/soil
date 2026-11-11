package symbol

import (
	"go/types"
	"golang.org/x/tools/go/packages"
)

func namedType(
	p *packages.Package,
	name string,
) *types.Named {
	o, okay := p.Types.Scope().Lookup(name).(*types.TypeName)

	if !okay || o.IsAlias() {
		return nil
	}

	n, okay := o.Type().(*types.Named)

	if !okay {
		return nil
	}

	return n
}
