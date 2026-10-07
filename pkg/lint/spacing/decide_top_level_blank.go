package spacing

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/types/spacing_shape"
	"strings"
)

func (s *Spacing) decideTopLevelBlank(h *spacing_shape.Shape) {
	pastIsVariable := strings.HasPrefix(h.PastTrimmed, "var ")
	pastIsConstant := strings.HasPrefix(h.PastTrimmed, "const ")
	sameKind := (pastIsVariable && h.Variable) || (pastIsConstant && h.Constant)

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
