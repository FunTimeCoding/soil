package error_answer

import (
	"fmt"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"slices"
)

func judge(
	p *packages.Package,
	results *output.Results,
	classes []string,
	body []ast.Stmt,
	position token.Pos,
) bool {
	if classes == nil {
		return false
	}

	observed := observe(p, body)

	if observed == nil {
		return slices.Contains(classes, errorsConstant.UnavailableClass)
	}

	for _, class := range classes {
		want := errorsConstant.Answers[class]

		if matches(want, observed) {
			continue
		}

		results.AddConcern(
			concern.NewPosition(
				constant.ErrorAnswerKey,
				fmt.Sprintf(
					"%s answers %s, this branch answers %s (%s)",
					class,
					describe(want, observed.Status == 0),
					describe(observed, observed.Status == 0),
					constant.ErrorAnswerReference,
				),
				p.Fset.Position(position),
			),
		)
	}

	return slices.Contains(classes, errorsConstant.UnavailableClass)
}
