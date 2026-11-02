package revision

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"
	"go/ast"
	"go/token"
)

func functionSymbol(
	s *token.FileSet,
	f *ast.File,
	d *ast.FuncDecl,
	q *qualifier,
) *symbol.Symbol {
	receiver := ""
	kind := constant.SymbolFunction

	if d.Recv != nil && len(d.Recv.List) > 0 {
		receiver = receiverName(d.Recv.List[0].Type)
		kind = constant.SymbolMethod
	}

	position := s.Position(d.Pos())
	result := symbol.New(
		packageDirectory(s, f),
		f.Name.Name,
		receiver,
		d.Name.Name,
		kind,
		position.Filename,
		position.Line,
	)
	result.Start, result.End = span(s, d)
	q.qualify(d, localNames(d))
	result.Body = render(s, d)
	result.Shape = anonymous(s, d, d.Name)
	result.Signature = render(
		s,
		&ast.FuncDecl{Recv: d.Recv, Name: d.Name, Type: d.Type},
	)

	return result
}
