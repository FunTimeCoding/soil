package decoration

import (
	"fmt"
	"github.com/dave/dst"
	"go/token"
)

func AddImport(
	file *dst.File,
	importPath string,
	alias string,
) {
	spec := &dst.ImportSpec{
		Path: &dst.BasicLit{
			Kind:  token.STRING,
			Value: fmt.Sprintf("%q", importPath),
		},
	}

	if alias != "" {
		spec.Name = dst.NewIdent(alias)
	}

	for _, d := range file.Decls {
		declaration, okay := d.(*dst.GenDecl)

		if !okay || declaration.Tok != token.IMPORT {
			continue
		}

		declaration.Lparen = true
		declaration.Rparen = true
		declaration.Specs = append(declaration.Specs, spec)

		return
	}

	declaration := &dst.GenDecl{Tok: token.IMPORT, Specs: []dst.Spec{spec}}
	declaration.Decs.After = dst.EmptyLine
	file.Decls = append([]dst.Decl{declaration}, file.Decls...)
}
