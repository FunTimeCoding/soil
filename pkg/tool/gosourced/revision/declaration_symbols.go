package revision

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"
	"go/ast"
	"go/token"
)

func declarationSymbols(
	s *token.FileSet,
	f *ast.File,
	d *ast.GenDecl,
	q *qualifier,
) []*symbol.Symbol {
	switch d.Tok {
	case token.TYPE:
		var result []*symbol.Symbol

		for _, spec := range d.Specs {
			if t, isType := spec.(*ast.TypeSpec); isType {
				result = append(result, typeSymbol(s, f, t, q))
			}
		}

		return result
	case token.CONST, token.VAR:
		return valueSymbols(s, f, d, q)
	default:
		return nil
	}
}
