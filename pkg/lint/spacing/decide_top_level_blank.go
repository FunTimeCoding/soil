package spacing

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func (s *Spacing) decideTopLevelBlank(h *shape) {
	pastIsVariable := strings.HasPrefix(h.pastTrimmed, "var ")
	pastIsConstant := strings.HasPrefix(h.pastTrimmed, "const ")
	sameKind := (pastIsVariable && h.variable) || (pastIsConstant && h.constant)

	if sameKind {
		s.concern(
			constant.ExtraneousTopLevelBlankKey,
			constant.ExtraneousTopLevelBlankText,
			s.pendingBlankLine,
			"",
		)

		return
	}

	s.report.ChangedLine("")
}
