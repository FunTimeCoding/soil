package decoration

import (
	"github.com/dave/dst"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func (s *Set) DecoratedNode(
	p *packages.Package,
	node ast.Node,
) dst.Node {
	dec, exists := s.Decorators[p]

	if !exists {
		return nil
	}

	return dec.Dst.Nodes[node]
}
