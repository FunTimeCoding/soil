package spacing

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/types/spacing_shape"
)

func (s *Spacing) extraneousBlank(h *spacing_shape.Shape) {
	s.concern(
		constant.ExtraneousBlankLineKey,
		constant.ExtraneousBlankLineText,
		h.Number,
		h.Line,
	)

	if !s.pendingBlank {
		s.needBlankAfterClosingBrace = false
	}
}
