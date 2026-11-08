package census

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/ast"
	value "go/constant"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func declaredTable(
	all []*packages.Package,
	m *types.Named,
) string {
	for _, p := range all {
		if p.PkgPath != m.Obj().Pkg().Path() {
			continue
		}

		for _, f := range p.Syntax {
			for _, d := range f.Decls {
				x, okay := d.(*ast.FuncDecl)

				if !okay || x.Body == nil || x.Name.Name != constant.TableNameMethod {
					continue
				}

				o, okay := p.TypesInfo.Defs[x.Name].(*types.Func)

				if !okay || owner(o) == nil || !sameType(owner(o), m) ||
					len(x.Body.List) != 1 {
					continue
				}

				r, okay := x.Body.List[0].(*ast.ReturnStmt)

				if !okay || len(r.Results) != 1 {
					continue
				}

				v := p.TypesInfo.Types[r.Results[0]].Value

				if v != nil && v.Kind() == value.String {
					return value.StringVal(v)
				}
			}
		}
	}

	return join.Empty(
		label(m),
		constant.MemberSeparator,
		constant.TableNameMethod,
		"()",
	)
}
