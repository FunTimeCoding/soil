package service

import (
	"go/ast"
	"go/token"
)

func expandFileQuerySymbols(file *ast.File) []*querySymbol {
	var result []*querySymbol

	for _, d := range file.Decls {
		switch declaration := d.(type) {
		case *ast.FuncDecl:
			receiver := receiverName(declaration)

			if declaration.Name.Name == "init" && receiver == "" {
				continue
			}

			result = append(
				result,
				&querySymbol{
					name:     declaration.Name.Name,
					receiver: receiver,
				},
			)
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
							&querySymbol{name: n.Name},
						)
					}
				case *ast.TypeSpec:
					result = append(
						result,
						&querySymbol{name: spec.Name.Name},
					)
				}
			}
		}
	}

	return result
}
