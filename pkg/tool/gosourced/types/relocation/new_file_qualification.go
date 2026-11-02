package relocation

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func NewFileQualification(
	file *ast.File,
	p *packages.Package,
	sourcePackagePath string,
) *FileQualification {
	return &FileQualification{
		File:        file,
		Owner:       p,
		SamePackage: p.PkgPath == sourcePackagePath,
		Idents:      make(map[*ast.Ident]string),
	}
}
