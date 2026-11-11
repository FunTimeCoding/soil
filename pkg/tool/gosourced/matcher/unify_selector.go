package matcher

import (
	"go/ast"
	"go/types"
)

func (m *Matcher) unifySelector(
	p *ast.SelectorExpr,
	s *ast.SelectorExpr,
) bool {
	qualifier, plain := p.X.(*ast.Ident)

	if !plain {
		return m.Unify(p.X, s.X) && m.Unify(p.Sel, s.Sel)
	}

	if _, isHole := m.holes[qualifier.Name]; isHole {
		return m.Unify(p.X, s.X) && m.Unify(p.Sel, s.Sel)
	}

	var name *types.PkgName

	if i, okay := s.X.(*ast.Ident); okay {
		name, _ = m.information.Uses[i].(*types.PkgName)
	}

	if path, resolved := m.qualifiers[qualifier.Name]; resolved {
		return name != nil && name.Imported().Path() == path &&
			m.Unify(p.Sel, s.Sel)
	}

	return m.Unify(p.X, s.X) && m.Unify(p.Sel, s.Sel)
}
