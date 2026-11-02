package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func releasedFields(
	p *packages.Package,
	methods map[string]*ast.FuncDecl,
) []fact.Release {
	return releasedFrom(p, methods, constant.CloseMethod)
}
