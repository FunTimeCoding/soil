package spacing

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func (s *Spacing) decideHeldBlank(h *shape) bool {
	if strings.HasPrefix(h.trimmed, constant.CommentPrefix) {
		s.report.ChangedLine("")
		s.report.ChangedLine(h.line)
		s.pendingBlank = false
		s.needBlankAfterClosingBrace = false

		return true
	}

	s.pendingBlank = false

	if h.topLevel {
		s.decideTopLevelBlank(h)

		return false
	}

	if h.pastOpensBlock {
		s.concern(
			constant.BlankInsideFunctionKey,
			constant.BlankInsideFunctionText,
			s.pendingBlankLine,
			"",
		)
	} else if s.needBlankAfterClosingBrace {
		s.report.ChangedLine("")
		s.needBlankAfterClosingBrace = false
	} else if h.controlStart || h.exit || h.deferral {
		s.report.ChangedLine("")
	} else {
		s.concern(
			constant.BlankInsideFunctionKey,
			constant.BlankInsideFunctionText,
			s.pendingBlankLine,
			"",
		)
	}

	return false
}
