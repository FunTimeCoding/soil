package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/literal"
	"go/ast"
	"go/types"
)

func pointerForm(site *literal.Site) bool {
	lit, okay := site.Target.(*ast.CompositeLit)

	if !okay || site.Outer != site.Target {
		return true
	}

	if lit.Type != nil {
		return false
	}

	_, pointer := site.Package.TypesInfo.TypeOf(lit).(*types.Pointer)

	return pointer
}
