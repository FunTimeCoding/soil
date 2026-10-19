package error_answer

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"net/http"
)

func checkBareUnavailable(
	p *packages.Package,
	results *output.Results,
	file *ast.File,
	guarded []ast.Node,
) {
	ast.Inspect(
		file,
		func(n ast.Node) bool {
			e, okay := n.(ast.Expr)

			if !okay || statusOf(p, e) != http.StatusServiceUnavailable {
				return true
			}

			if captures(e) || within(guarded, e) {
				return false
			}

			results.AddConcern(
				concern.NewPosition(
					constant.ErrorAnswerKey,
					fmt.Sprintf(
						"an uncaptured 503 answers only unavailable; a transport failure answers 503 ErrorResponse (%s)",
						constant.ErrorAnswerReference,
					),
					p.Fset.Position(e.Pos()),
				),
			)

			return false
		},
	)
}
