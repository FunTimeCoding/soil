package stray_variable

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func isBuildMetadata(
	p *packages.Package,
	name *ast.Ident,
) bool {
	return p.Types.Name() == "main" &&
		constant.BuildMetadataVariables[name.Name]
}
