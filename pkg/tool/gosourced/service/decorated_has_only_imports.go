package service

import (
	"github.com/dave/dst"
	"go/token"
)

func decoratedHasOnlyImports(file *dst.File) bool {
	for _, d := range file.Decls {
		g, okay := d.(*dst.GenDecl)

		if !okay || g.Tok != token.IMPORT {
			return false
		}
	}

	return true
}
