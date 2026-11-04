package naming

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func NewSite(
	owner *packages.Package,
	file *ast.File,
	signature *ast.FuncType,
	field int,
	index int,
	label string,
) *Site {
	return &Site{
		Owner:     owner,
		File:      file,
		Signature: signature,
		Field:     field,
		Index:     index,
		Label:     label,
	}
}
