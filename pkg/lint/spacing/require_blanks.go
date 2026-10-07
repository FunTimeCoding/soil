package spacing

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/types/spacing_shape"
	"strings"
)

func (s *Spacing) requireBlanks(h *spacing_shape.Shape) {
	preceded := !s.pastWasBlank && s.pastLine != ""

	if h.ControlStart && preceded && !h.PastOpensBlock {
		s.report.ChangedLine("")
		s.concern(
			constant.MissingBlankBeforeControlKey,
			constant.MissingBlankBeforeControlText,
			h.Number,
			h.Line,
		)
		s.needBlankAfterClosingBrace = false
	}

	if h.Exit && preceded && !h.PastOpensBlock {
		s.report.ChangedLine("")
		s.concern(
			constant.MissingBlankBeforeExitKey,
			constant.MissingBlankBeforeExitText,
			h.Number,
			h.Line,
		)
		s.needBlankAfterClosingBrace = false
	}

	if h.TopLevelDeclaration && preceded && !h.PastOpensBlock {
		s.report.ChangedLine("")
		s.concern(
			constant.MissingBlankBeforeDeclarationKey,
			constant.MissingBlankBeforeDeclarationText,
			h.Number,
			h.Line,
		)
	}

	pastIsVariable := strings.HasPrefix(h.PastTrimmed, "var ")
	pastIsConstant := strings.HasPrefix(h.PastTrimmed, "const ")
	crossKind := (pastIsVariable && h.Constant) || (pastIsConstant && h.Variable)

	if crossKind && preceded {
		s.report.ChangedLine("")
		s.concern(
			constant.MissingBlankBetweenVariableConstantKey,
			constant.MissingBlankBetweenVariableConstantText,
			h.Number,
			h.Line,
		)
	}

	if s.needBlankAfterClosingBrace && !h.Blank && !h.ClosingBrace {
		s.report.ChangedLine("")
		s.concern(
			constant.MissingBlankAfterControlKey,
			constant.MissingBlankAfterControlText,
			h.Number,
			h.Line,
		)
	}
}
