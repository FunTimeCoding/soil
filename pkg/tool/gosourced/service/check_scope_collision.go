package service

import (
	"fmt"
	"golang.org/x/tools/go/packages"
)

func checkScopeCollision(
	p *packages.Package,
	targetName string,
) error {
	if p.Types.Scope().Lookup(targetName) != nil {
		return fmt.Errorf("%s already exists in %s", targetName, p.PkgPath)
	}

	return nil
}
