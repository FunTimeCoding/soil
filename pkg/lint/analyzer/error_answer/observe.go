package error_answer

import (
	"github.com/funtimecoding/soil/pkg/errors/types/answer"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func observe(
	p *packages.Package,
	body []ast.Stmt,
) *answer.Answer {
	for _, s := range body {
		r, okay := s.(*ast.ReturnStmt)

		if !okay {
			continue
		}

		if len(r.Results) == 0 {
			return nil
		}

		return answerOf(p, r.Results[0])
	}

	return nil
}
