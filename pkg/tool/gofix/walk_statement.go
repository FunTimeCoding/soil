package gofix

import "github.com/dave/dst"

func walkStatement(
	t dst.Stmt,
	parentLit *dst.CompositeLit,
	walk func(
		dst.Node,
		*dst.CompositeLit,
		int,
	),
) {
	switch s := t.(type) {
	case *dst.ExprStmt:
		walk(s.X, parentLit, 0)
	case *dst.AssignStmt:
		for _, rhs := range s.Rhs {
			walk(rhs, parentLit, 0)
		}
	case *dst.ReturnStmt:
		for _, result := range s.Results {
			walk(result, parentLit, 0)
		}
	case *dst.IfStmt:
		if s.Init != nil {
			walkStatement(s.Init, parentLit, walk)
		}

		walkBlock(s.Body, parentLit, walk)

		if s.Else != nil {
			walkStatement(s.Else, parentLit, walk)
		}
	case *dst.ForStmt:
		walkBlock(s.Body, parentLit, walk)
	case *dst.RangeStmt:
		walkBlock(s.Body, parentLit, walk)
	case *dst.SwitchStmt:
		walkBlock(s.Body, parentLit, walk)
	case *dst.TypeSwitchStmt:
		walkBlock(s.Body, parentLit, walk)
	case *dst.SelectStmt:
		walkBlock(s.Body, parentLit, walk)
	case *dst.CaseClause:
		for _, t := range s.Body {
			walkStatement(t, parentLit, walk)
		}
	case *dst.CommClause:
		for _, t := range s.Body {
			walkStatement(t, parentLit, walk)
		}
	case *dst.LabeledStmt:
		walkStatement(s.Stmt, parentLit, walk)
	case *dst.BlockStmt:
		walkBlock(s, parentLit, walk)
	case *dst.DeferStmt:
		walk(s.Call, parentLit, 0)
	case *dst.GoStmt:
		walk(s.Call, parentLit, 0)
	case *dst.SendStmt:
		walk(s.Value, parentLit, 0)
	case *dst.DeclStmt:
		declaration, okay := s.Decl.(*dst.GenDecl)

		if !okay {
			return
		}

		longest := longestSpecName(declaration)

		for _, spec := range declaration.Specs {
			v, okay := spec.(*dst.ValueSpec)

			if !okay {
				continue
			}

			padding := 0

			if len(v.Names) > 0 && longest > len(v.Names[0].Name) {
				padding = longest - len(v.Names[0].Name)
			}

			for _, value := range v.Values {
				walk(value, parentLit, padding)
			}
		}
	}
}
