package policy

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"mvdan.cc/sh/v3/syntax"
)

func (w *Walk) statement(s *syntax.Stmt) bool {
	result := true

	if s.Background || s.Coprocess || s.Disown {
		result = w.block(constant.BlockerBackground)
	}

	for _, r := range s.Redirs {
		if !w.redirect(r) {
			result = false
		}
	}

	return w.command(s.Cmd) && result
}
