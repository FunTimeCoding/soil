package error_answer

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"strings"
)

func checkPredicate(
	p *packages.Package,
	results *output.Results,
	f *ast.FuncDecl,
) {
	if f.Body == nil || len(f.Body.List) != 1 {
		return
	}

	r, okay := f.Body.List[0].(*ast.ReturnStmt)

	if !okay || len(r.Results) != 1 {
		return
	}

	classes := conditionClasses(p, r.Results[0])

	if len(classes) == 0 {
		return
	}

	results.AddConcern(
		concern.NewPosition(
			constant.ErrorAnswerKey,
			fmt.Sprintf(
				"%s hides %s from the answer; branch on each class where the answer is built (%s)",
				f.Name.Name,
				strings.Join(classes, ", "),
				constant.ErrorAnswerReference,
			),
			p.Fset.Position(f.Pos()),
		),
	)
}
