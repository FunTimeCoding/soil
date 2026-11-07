package policy

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func (w *Walk) copyOperands(arguments []*syntax.Word) bool {
	list, options := operands(arguments)

	for _, o := range options {
		if o == constant.TargetOption ||
			strings.HasPrefix(o, constant.TargetOptionLong) {
			return false
		}
	}

	if len(list) < 2 {
		return false
	}

	return w.scratch(list[len(list)-1])
}
