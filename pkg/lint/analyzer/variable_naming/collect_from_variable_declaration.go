package variable_naming

import (
	"go/ast"
	"go/token"
	"go/types"
)

func collectFromVariableDeclaration(
	y *types.Info,
	s *ast.DeclStmt,
	result *[]TypedVariable,
) {
	g, okay := s.Decl.(*ast.GenDecl)

	if !okay || g.Tok != token.VAR {
		return
	}

	for _, spec := range g.Specs {
		v, okay := spec.(*ast.ValueSpec)

		if !okay {
			continue
		}

		for _, name := range v.Names {
			if name.Name == "_" {
				continue
			}

			o := y.ObjectOf(name)

			if o == nil {
				continue
			}

			*result = append(
				*result,
				TypedVariable{
					ident:       name,
					typ:         o.Type(),
					precedence:  typePrecedence(o.Type()),
					scopedNames: collectScopedNames(o),
				},
			)
		}
	}
}
