package service

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func checkMethodCollision(
	p *packages.Package,
	targetName string,
	receiver string,
) error {
	o := p.Types.Scope().Lookup(receiver)

	if o == nil {
		return nil
	}

	named, okay := o.Type().(*types.Named)

	if !okay {
		return nil
	}

	for i := range named.NumMethods() {
		if named.Method(i).Name() == targetName {
			return conflict.Format(
				"method %s already exists on %s in %s",
				targetName,
				receiver,
				p.PkgPath,
			)
		}
	}

	return nil
}
