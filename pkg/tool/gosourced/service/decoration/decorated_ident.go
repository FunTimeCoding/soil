package decoration

import (
	"github.com/dave/dst"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func (s *Set) DecoratedIdent(
	p *packages.Package,
	ident *ast.Ident,
) *dst.Ident {
	dec, exists := s.Decorators[p]

	if !exists {
		return nil
	}

	result, _ := dec.Dst.Nodes[ident].(*dst.Ident)

	return result
}
