package error_answer

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func checkFile(
	p *packages.Package,
	results *output.Results,
	file *ast.File,
) {
	var guarded []ast.Node
	ast.Inspect(
		file,
		func(n ast.Node) bool {
			if x, okay := n.(*ast.FuncDecl); okay {
				checkPredicate(p, results, x)
			}

			if x, okay := n.(*ast.IfStmt); okay &&
				judge(p, results, ifClasses(p, x), x.Body.List, x.Cond.Pos()) {
				guarded = append(guarded, x.Body)
			}

			if x, okay := n.(*ast.CaseClause); okay && len(x.List) > 0 &&
				judge(p, results, listClasses(p, x.List), x.Body, x.Pos()) {
				guarded = append(guarded, x)
			}

			return true
		},
	)
	checkBareUnavailable(p, results, file, guarded)
}
