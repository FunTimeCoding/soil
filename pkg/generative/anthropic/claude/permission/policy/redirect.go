package policy

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func (w *Walk) redirect(r *syntax.Redirect) bool {
	result := r.Word == nil || w.nested(r.Word)

	if r.Hdoc != nil && !w.nested(r.Hdoc) {
		result = false
	}

	switch r.Op {
	case syntax.RdrIn, syntax.Hdoc, syntax.DashHdoc, syntax.WordHdoc:
		return result
	case syntax.DplIn, syntax.DplOut:
		if v, okay := literal(r.Word); okay &&
			(v == "-" || strings.Trim(v, "0123456789") == "") {
			return result
		}

		return (w.writable(r.Word) || w.block(constant.BlockerRedirect)) &&
			result
	case syntax.RdrOut, syntax.AppOut, syntax.RdrAll, syntax.AppAll,
		syntax.RdrClob, syntax.RdrInOut:

		return (w.writable(r.Word) || w.block(constant.BlockerRedirect)) &&
			result
	default:
		return w.block(constant.BlockerRedirect)
	}
}
