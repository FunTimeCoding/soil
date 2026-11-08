package census

import (
	"go/types"
	"golang.org/x/tools/go/packages"
)

func censusType(
	p *packages.Package,
	name string,
) *types.Named {
	o, okay := p.Types.Scope().Lookup(name).(*types.TypeName)

	if !okay {
		return nil
	}

	n, okay := o.Type().(*types.Named)

	if !okay {
		return nil
	}

	if _, okay := n.Underlying().(*types.Struct); !okay {
		return nil
	}

	return n
}
