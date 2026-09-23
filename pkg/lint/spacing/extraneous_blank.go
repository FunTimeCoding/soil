package spacing

import "github.com/funtimecoding/soil/pkg/lint/constant"

func (s *Spacing) extraneousBlank(h *shape) {
	s.concern(
		constant.ExtraneousBlankLineKey,
		constant.ExtraneousBlankLineText,
		h.number,
		h.line,
	)

	if !s.pendingBlank {
		s.needBlankAfterClosingBrace = false
	}
}
