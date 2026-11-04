package symbol

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/naming"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func namingSites(
	all []*packages.Package,
	set *token.FileSet,
	scope map[string]bool,
	request *naming.Request,
	names map[string][]string,
	explicit map[string]string,
	index map[string]*types.Package,
	r *output.Results,
) []*naming.Site {
	var result []*naming.Site
	seen := map[string]bool{}
	listed := func(f *ast.FuncDecl) bool {
		if len(request.Functions) == 0 {
			return true
		}

		for _, l := range request.Functions {
			if l.Name == f.Name.Name && l.Receiver == receiverName(f) {
				return true
			}
		}

		return false
	}

	for _, p := range all {
		if !scope[p.PkgPath] {
			continue
		}

		for _, file := range p.Syntax {
			filename := set.Position(file.Pos()).Filename

			if seen[filename] {
				continue
			}

			seen[filename] = true
			t, _, e := evaluate(
				set,
				p,
				file.Name.Pos(),
				request.Type,
				true,
				names,
				explicit,
				index,
			)

			if e != nil {
				r.AddConcern(
					concern.NewFile(
						constant.ConcernValidation,
						e.Error(),
						filename,
						false,
					),
				)

				return nil
			}

			visit := func(
				signature *ast.FuncType,
				body *ast.BlockStmt,
				label string,
			) {
				site := namingSite(
					p,
					file,
					set,
					signature,
					body,
					t,
					request.Name,
					label,
					r,
				)

				if site != nil {
					result = append(result, site)
				}
			}
			literals := func(root ast.Node, label string) {
				ast.Inspect(
					root,
					func(n ast.Node) bool {
						if l, okay := n.(*ast.FuncLit); okay {
							visit(l.Type, l.Body, label)
						}

						return true
					},
				)
			}

			for _, d := range file.Decls {
				switch v := d.(type) {
				case *ast.FuncDecl:
					if v.Body == nil || !listed(v) {
						continue
					}

					visit(v.Type, v.Body, v.Name.Name)
					literals(v.Body, v.Name.Name)
				case *ast.GenDecl:
					if len(request.Functions) == 0 {
						literals(v, "literal")
					}
				}
			}
		}
	}

	return result
}
