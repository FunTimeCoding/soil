package census

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	"strings"
)

func countCallers(
	loaded []*packages.Package,
	home string,
	subject *types.Named,
	module string,
	methods map[string]*reading,
) {
	seen := map[string]bool{}

	for _, p := range loaded {
		for _, f := range p.Syntax {
			file := p.Fset.Position(f.Pos()).Filename
			ast.Inspect(
				f,
				func(n ast.Node) bool {
					i, okay := n.(*ast.Ident)

					if !okay {
						return true
					}

					o, okay := p.TypesInfo.Uses[i].(*types.Func)

					if !okay || owner(o) == nil || !sameType(owner(o), subject) {
						return true
					}

					target := methods[o.Name()]
					at := p.Fset.Position(i.Pos()).String()

					if target == nil || seen[at] {
						return true
					}

					seen[at] = true
					path := strings.TrimSuffix(
						p.PkgPath,
						constant.TestClauseSuffix,
					)
					k := callerKey{
						kind:        callerKind(path, home, file),
						packagePath: relative(module, path),
						function:    enclosing(f, i.Pos()),
					}
					target.callers[k]++

					return true
				},
			)
		}
	}
}
