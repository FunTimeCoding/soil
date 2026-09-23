package spacing

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func (s *Spacing) requireBlanks(h *shape) {
	preceded := !s.pastWasBlank && s.pastLine != ""

	if h.controlStart && preceded && !h.pastOpensBlock {
		s.report.ChangedLine("")
		s.concern(
			constant.MissingBlankBeforeControlKey,
			constant.MissingBlankBeforeControlText,
			h.number,
			h.line,
		)
		s.needBlankAfterClosingBrace = false
	}

	if h.exit && preceded && !h.pastOpensBlock {
		s.report.ChangedLine("")
		s.concern(
			constant.MissingBlankBeforeExitKey,
			constant.MissingBlankBeforeExitText,
			h.number,
			h.line,
		)
		s.needBlankAfterClosingBrace = false
	}

	if h.topLevelDeclaration && preceded && !h.pastOpensBlock {
		s.report.ChangedLine("")
		s.concern(
			constant.MissingBlankBeforeDeclarationKey,
			constant.MissingBlankBeforeDeclarationText,
			h.number,
			h.line,
		)
	}

	pastIsVariable := strings.HasPrefix(h.pastTrimmed, "var ")
	pastIsConstant := strings.HasPrefix(h.pastTrimmed, "const ")
	crossKind := (pastIsVariable && h.constant) || (pastIsConstant && h.variable)

	if crossKind && preceded {
		s.report.ChangedLine("")
		s.concern(
			constant.MissingBlankBetweenVariableConstantKey,
			constant.MissingBlankBetweenVariableConstantText,
			h.number,
			h.line,
		)
	}

	if s.needBlankAfterClosingBrace && !h.blank && !h.closingBrace {
		s.report.ChangedLine("")
		s.concern(
			constant.MissingBlankAfterControlKey,
			constant.MissingBlankAfterControlText,
			h.number,
			h.line,
		)
	}
}
