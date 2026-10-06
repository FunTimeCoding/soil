package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/pattern_site"
	"go/ast"
	"go/token"
)

func expandFileQuerySymbols(file *ast.File) []*pattern_site.QuerySymbol {
	var result []*pattern_site.QuerySymbol

	for _, d := range file.Decls {
		switch declaration := d.(type) {
		case *ast.FuncDecl:
			receiver := receiverName(declaration)

			if declaration.Name.Name == "init" && receiver == "" {
				continue
			}

			result = append(
				result,
				pattern_site.NewQuerySymbol(declaration.Name.Name, receiver))
		case *ast.GenDecl:
			if declaration.Tok == token.IMPORT {
				continue
			}

			for _, s := range declaration.Specs {
				switch spec := s.(type) {
				case *ast.ValueSpec:
					for _, n := range spec.Names {
						if n.Name == "_" {
							continue
						}

						result = append(
							result,
							pattern_site.NewQuerySymbol(n.Name, ""),
						)
					}
				case *ast.TypeSpec:
					result = append(
						result,
						pattern_site.NewQuerySymbol(spec.Name.Name, ""),
					)
				}
			}
		}
	}

	return result
}
