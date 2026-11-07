package policy

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"mvdan.cc/sh/v3/syntax"
)

func (w *Walk) writable(word *syntax.Word) bool {
	if v, okay := w.resolve(word); okay && v == constant.NullDevice {
		return true
	}

	return w.source(w.scratch(word), constant.ScratchSource)
}
