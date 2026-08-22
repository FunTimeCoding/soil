package service

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"golang.org/x/tools/go/packages"
)

func checkScopeCollision(
	p *packages.Package,
	targetName string,
) error {
	if p.Types.Scope().Lookup(targetName) != nil {
		return conflict.Format("%s already exists in %s", targetName, p.PkgPath)
	}

	return nil
}
