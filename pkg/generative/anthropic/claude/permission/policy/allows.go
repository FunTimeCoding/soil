package policy

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func (w *Walk) Allows(command string) bool {
	f, e := syntax.NewParser().Parse(strings.NewReader(command), "")

	if e != nil {
		return w.block(constant.BlockerUnparsed)
	}

	if !w.statements(f.Stmts) {
		return false
	}

	if w.commands == 0 {
		return w.block(constant.BlockerEmpty)
	}

	return true
}
