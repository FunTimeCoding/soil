package spacing

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/types/spacing_shape"
	"strings"
)

func (s *Spacing) decideHeldBlank(h *spacing_shape.Shape) bool {
	if strings.HasPrefix(h.Trimmed, constant.CommentPrefix) {
		s.report.ChangedLine("")
		s.report.ChangedLine(h.Line)
		s.pendingBlank = false
		s.needBlankAfterClosingBrace = false

		return true
	}

	s.pendingBlank = false

	if h.TopLevel {
		s.decideTopLevelBlank(h)

		return false
	}

	if h.PastOpensBlock {
		s.concern(
			constant.BlankInsideFunctionKey,
			constant.BlankInsideFunctionText,
			s.pendingBlankLine,
			"",
		)
	} else if s.needBlankAfterClosingBrace {
		s.report.ChangedLine("")
		s.needBlankAfterClosingBrace = false
	} else if h.ControlStart || h.Exit || h.Deferral {
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
