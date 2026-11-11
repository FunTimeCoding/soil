package relocator

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/relocation"
	"golang.org/x/tools/go/packages"
)

func pathPackage(
	plan *relocation.Plan,
	path string,
) *packages.Package {
	if path == plan.Source.PkgPath {
		return plan.Source
	}

	if p, okay := plan.Source.Imports[path]; okay {
		return p
	}

	if plan.Target == nil {
		return nil
	}

	return plan.Target.Imports[path]
}
