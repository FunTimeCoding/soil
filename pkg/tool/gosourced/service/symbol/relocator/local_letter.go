package relocator

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/relocation"
	"go/ast"
	"go/types"
	"unicode"
)

func localLetter(
	q *relocation.FileQualification,
	v *types.Var,
	held func(string) bool,
) string {
	scope := v.Parent()

	for _, r := range v.Name() {
		if r > unicode.MaxASCII || !unicode.IsLetter(r) {
			continue
		}

		letter := string(unicode.ToLower(r))

		if held(letter) {
			continue
		}

		used := false
		ast.Inspect(
			q.File,
			func(n ast.Node) bool {
				if i, okay := n.(*ast.Ident); okay && i.Name == letter &&
					i.Pos() >= scope.Pos() && i.Pos() < scope.End() {
					used = true
				}

				return !used
			},
		)

		if !used {
			return letter
		}
	}

	return ""
}
