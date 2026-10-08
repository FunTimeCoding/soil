package gofix

import (
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/element_format"
	"go/ast"
	"go/token"
)

func walkFormatEdits(
	destinationFile *dst.File,
	dec *decorator.Decorator,
	fileSet *token.FileSet,
	source []byte,
	collapse bool,
) []*FormatChange {
	var result []*FormatChange
	stale := 0
	var walk func(
		dst.Node,
		*dst.CompositeLit,
		int,
	)
	walk = func(
		n dst.Node,
		parentLit *dst.CompositeLit,
		extraPadding int,
	) {
		switch node := n.(type) {
		case *dst.CallExpr:
			changed := false

			if stale == 0 && len(node.Args) > 0 {
				astNode := dec.Ast.Nodes[node]
				astCall, okay := astNode.(*ast.CallExpr)

				if okay {
					el := element_format.FromCall(astCall)
					el.HasComments = len(node.Decs.Lparen) > 0
					el.Padding = callFieldPadding(node, parentLit)

					if el.Padding == 0 {
						el.Padding = extraPadding
					}

					if element_format.Apply(
						fileSet,
						el,
						source,
						node.Args,
						collapse,
					) {
						changed = true
						result = append(
							result,
							&FormatChange{
								Kind:    "call_format",
								Message: "formatted call",
								Offset:  fileSet.Position(astCall.Lparen).Offset,
							},
						)
					}
				}
			}

			if changed {
				stale++
			}

			walk(node.Fun, parentLit, 0)

			for _, arg := range node.Args {
				walk(arg, parentLit, 0)
			}

			if changed {
				stale--
			}
		case *dst.CompositeLit:
			changed := false

			if stale == 0 && len(node.Elts) > 0 {
				astNode := dec.Ast.Nodes[node]
				astLit, okay := astNode.(*ast.CompositeLit)

				if okay {
					el := element_format.FromLiteral(astLit)
					el.Padding = structFieldPadding(node, parentLit)
					el.HasComments = len(node.Decs.Lbrace) > 0

					if el.Padding == 0 {
						el.Padding = extraPadding
					}

					if element_format.Apply(
						fileSet,
						el,
						source,
						node.Elts,
						collapse,
					) {
						changed = true
						result = append(
							result,
							&FormatChange{
								Kind:    "composite_format",
								Message: "formatted composite literal",
								Offset:  fileSet.Position(astLit.Lbrace).Offset,
							},
						)
					}
				}
			}

			if changed {
				stale++
			}

			if node.Type != nil {
				walk(node.Type, parentLit, 0)
			}

			for _, el := range node.Elts {
				walk(el, node, 0)
			}

			if changed {
				stale--
			}
		case *dst.KeyValueExpr:
			walk(node.Key, parentLit, 0)
			walk(node.Value, parentLit, 0)
		case *dst.UnaryExpr:
			walk(node.X, parentLit, 0)
		case *dst.SelectorExpr:
			walk(node.X, parentLit, 0)
		case *dst.IndexExpr:
			walk(node.X, parentLit, 0)
			walk(node.Index, parentLit, 0)
		case *dst.SliceExpr:
			walk(node.X, parentLit, 0)
		case *dst.StarExpr:
			walk(node.X, parentLit, 0)
		case *dst.ParenExpr:
			walk(node.X, parentLit, 0)
		case *dst.FuncLit:
			walkBlock(node.Body, parentLit, walk)
		}
	}
	dst.Inspect(
		destinationFile,
		func(n dst.Node) bool {
			switch node := n.(type) {
			case *dst.FuncDecl:
				if node.Body != nil {
					walkBlock(node.Body, nil, walk)
				}

				return false
			case *dst.GenDecl:
				longest := longestSpecName(node)

				for _, spec := range node.Specs {
					v, okay := spec.(*dst.ValueSpec)

					if !okay {
						continue
					}

					padding := 0

					if len(v.Names) > 0 && longest > len(v.Names[0].Name) {
						padding = longest - len(v.Names[0].Name)
					}

					for _, value := range v.Values {
						walk(value, nil, padding)
					}
				}

				return false
			}

			return true
		},
	)

	return result
}
