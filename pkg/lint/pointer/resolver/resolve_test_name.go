package resolver

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"strings"
)

func (r *Resolver) resolveTestName(span string) *pointer.Resolution {
	if r.Tests(strings.TrimSuffix(span, constant.CallSuffix)) {
		return pointer.NewLive("")
	}

	return pointer.NewUnchecked(constant.ReasonTest)
}
