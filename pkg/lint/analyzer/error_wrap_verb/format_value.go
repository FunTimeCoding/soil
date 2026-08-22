package error_wrap_verb

import (
	"go/ast"
	"go/constant"
	"golang.org/x/tools/go/packages"
)

func formatValue(
	p *packages.Package,
	argument ast.Expr,
) (string, bool) {
	t, okay := p.TypesInfo.Types[argument]

	if !okay || t.Value == nil || t.Value.Kind() != constant.String {
		return "", false
	}

	return constant.StringVal(t.Value), true
}
