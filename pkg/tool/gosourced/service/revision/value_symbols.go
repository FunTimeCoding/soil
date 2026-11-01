package revision

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"
	"go/ast"
	"go/token"
	"strconv"
)

func valueSymbols(
	s *token.FileSet,
	f *ast.File,
	d *ast.GenDecl,
	q *qualifier,
) []*symbol.Symbol {
	kind := constant.SymbolVariable

	if d.Tok == token.CONST {
		kind = constant.SymbolConstant
	}

	counted := d.Tok == token.CONST && usesIota(d)
	var inherited *ast.ValueSpec
	var result []*symbol.Symbol

	for index, spec := range d.Specs {
		v, isValue := spec.(*ast.ValueSpec)

		if !isValue {
			continue
		}

		q.qualify(v, localNames(v))
		rendered := &ast.ValueSpec{
			Names:  v.Names,
			Type:   v.Type,
			Values: v.Values,
		}

		if len(v.Values) > 0 {
			inherited = v
		} else if d.Tok == token.CONST && inherited != nil {
			rendered.Type = inherited.Type
			rendered.Values = inherited.Values
		}

		declaration := &ast.GenDecl{Tok: d.Tok, Specs: []ast.Spec{rendered}}
		body := render(s, declaration)
		shape := anonymous(s, declaration, v.Names...)

		if counted {
			position := join.Empty(constant.IotaMarker, strconv.Itoa(index))
			body = join.Empty(body, position)
			shape = join.Empty(shape, position)
		}

		result = append(result, specSymbols(s, f, v, kind, body, shape)...)
	}

	return result
}
