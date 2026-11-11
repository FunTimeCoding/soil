package symbol

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/ast"
	"go/token"
)

func expandFileSymbols(file *ast.File) ([]string, error) {
	var result []string

	for _, d := range file.Decls {
		switch declaration := d.(type) {
		case *ast.FuncDecl:
			if declaration.Recv != nil {
				return nil, validation.New(
					"file contains method %s - move it with %s, or its type with %s",
					declaration.Name.Name,
					constant.MoveMethods,
					constant.ExtractType,
				)
			}

			result = append(result, declaration.Name.Name)
		case *ast.GenDecl:
			if declaration.Tok == token.IMPORT {
				continue
			}

			for _, s := range declaration.Specs {
				switch spec := s.(type) {
				case *ast.ValueSpec:
					for _, n := range spec.Names {
						result = append(result, n.Name)
					}
				case *ast.TypeSpec:
					result = append(result, spec.Name.Name)
				}
			}
		}
	}

	if len(result) == 0 {
		return nil, validation.New("file declares no movable symbols")
	}

	return result, nil
}
