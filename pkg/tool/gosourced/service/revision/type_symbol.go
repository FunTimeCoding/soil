package revision

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"
	"go/ast"
	"go/token"
)

func typeSymbol(
	s *token.FileSet,
	f *ast.File,
	t *ast.TypeSpec,
	q *qualifier,
) *symbol.Symbol {
	position := s.Position(t.Pos())
	result := symbol.New(
		packageDirectory(s, f),
		f.Name.Name,
		"",
		t.Name.Name,
		constant.SymbolType,
		position.Filename,
		position.Line,
	)
	q.qualify(t, localNames(t))
	declaration := &ast.GenDecl{Tok: token.TYPE, Specs: []ast.Spec{t}}
	result.Body = render(s, declaration)
	result.Shape = anonymous(s, declaration, t.Name)

	return result
}
