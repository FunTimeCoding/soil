package service

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func newFileQualification(
	file *ast.File,
	p *packages.Package,
	sourcePackagePath string,
) *fileQualification {
	return &fileQualification{
		file:        file,
		samePackage: p.PkgPath == sourcePackagePath,
		idents:      make(map[*ast.Ident]string),
	}
}
