package match

import (
	"go/ast"
	"go/token"
)

func (m *Matcher) Unify(pattern ast.Node, site ast.Node) bool {
	switch p := pattern.(type) {
	case *ast.Ident:
		if _, isHole := m.holes[p.Name]; isHole {
			return m.holeMatch(p.Name, site)
		}

		s, okay := site.(*ast.Ident)

		if !okay {
			return false
		}

		if p.Name == m.querySymbol {
			return m.isAnchor(m.information.Uses[s])
		}

		return p.Name == s.Name
	case *ast.SelectorExpr:
		s, okay := site.(*ast.SelectorExpr)

		return okay && m.Unify(p.X, s.X) && m.Unify(p.Sel, s.Sel)
	case *ast.CallExpr:
		s, okay := site.(*ast.CallExpr)

		if !okay {
			return false
		}

		if !m.Unify(p.Fun, s.Fun) {
			return false
		}

		if len(p.Args) == 1 && p.Ellipsis.IsValid() {
			if hole, wildcard := p.Args[0].(*ast.Ident); wildcard {
				if _, isHole := m.holes[hole.Name]; isHole {
					return true
				}
			}
		}

		if len(p.Args) != len(s.Args) {
			return false
		}

		if (p.Ellipsis == token.NoPos) != (s.Ellipsis == token.NoPos) {
			return false
		}

		for i, argument := range p.Args {
			if !m.Unify(argument, s.Args[i]) {
				return false
			}
		}

		return true
	case *ast.BlockStmt:
		return true
	case *ast.IfStmt:
		s, okay := site.(*ast.IfStmt)

		if !okay {
			return false
		}

		if (p.Init == nil) != (s.Init == nil) {
			return false
		}

		if p.Init != nil && !m.Unify(p.Init, s.Init) {
			return false
		}

		if !m.Unify(p.Cond, s.Cond) {
			return false
		}

		if p.Else == nil {
			return true
		}

		return m.Unify(p.Else, s.Else)
	case *ast.ExprStmt:
		s, okay := site.(*ast.ExprStmt)

		return okay && m.Unify(p.X, s.X)
	case *ast.AssignStmt:
		s, okay := site.(*ast.AssignStmt)

		if !okay || p.Tok != s.Tok {
			return false
		}

		if len(p.Lhs) != len(s.Lhs) || len(p.Rhs) != len(s.Rhs) {
			return false
		}

		for i, left := range p.Lhs {
			if !m.Unify(left, s.Lhs[i]) {
				return false
			}
		}

		for i, right := range p.Rhs {
			if !m.Unify(right, s.Rhs[i]) {
				return false
			}
		}

		return true
	case *ast.ReturnStmt:
		s, okay := site.(*ast.ReturnStmt)

		if !okay || len(p.Results) != len(s.Results) {
			return false
		}

		for i, value := range p.Results {
			if !m.Unify(value, s.Results[i]) {
				return false
			}
		}

		return true
	case *ast.BasicLit:
		s, okay := site.(*ast.BasicLit)

		return okay && p.Kind == s.Kind && p.Value == s.Value
	case *ast.BinaryExpr:
		s, okay := site.(*ast.BinaryExpr)

		return okay && p.Op == s.Op &&
			m.Unify(p.X, s.X) &&
			m.Unify(p.Y, s.Y)
	case *ast.UnaryExpr:
		s, okay := site.(*ast.UnaryExpr)

		return okay && p.Op == s.Op && m.Unify(p.X, s.X)
	case *ast.ParenExpr:
		s, okay := site.(*ast.ParenExpr)

		return okay && m.Unify(p.X, s.X)
	case *ast.StarExpr:
		s, okay := site.(*ast.StarExpr)

		return okay && m.Unify(p.X, s.X)
	case *ast.DeferStmt:
		s, okay := site.(*ast.DeferStmt)

		return okay && m.Unify(p.Call, s.Call)
	case *ast.GoStmt:
		s, okay := site.(*ast.GoStmt)

		return okay && m.Unify(p.Call, s.Call)
	default:
		return false
	}
}
