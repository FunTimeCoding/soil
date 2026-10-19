package error_answer

import (
	"github.com/funtimecoding/soil/pkg/errors/types/answer"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"strings"
)

func answerOf(
	p *packages.Package,
	e ast.Expr,
) *answer.Answer {
	if status := statusOf(p, e); status != 0 {
		if captures(e) {
			return answer.NewCaptured(status)
		}

		return answer.NewUncaptured(status)
	}

	call, okay := e.(*ast.CallExpr)

	if !okay {
		return nil
	}

	f := callee(p, call)

	if f == nil || f.Pkg() == nil {
		return nil
	}

	if f.Name() == constant.ResponseFail &&
		strings.HasSuffix(f.Pkg().Path(), constant.ResponsePackageSuffix) {
		return answer.NewUncaptured(0)
	}

	if f.Name() == constant.CaptureFailMethod {
		return answer.NewCaptured(0)
	}

	return nil
}
