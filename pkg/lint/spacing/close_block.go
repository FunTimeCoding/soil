package spacing

import (
	"github.com/funtimecoding/soil/pkg/lint/types/spacing_shape"
	"strings"
)

func (s *Spacing) closeBlock(h *spacing_shape.Shape) {
	if !strings.HasPrefix(
		h.Trimmed,
		"}",
	) || h.ElseContinuation || h.EndsWithBrace {
		s.needBlankAfterClosingBrace = false

		return
	}

	isControl := false

	if len(s.blockStack) > 0 {
		isControl = s.blockStack[len(s.blockStack)-1]
		s.blockStack = s.blockStack[:len(s.blockStack)-1]
	}

	s.needBlankAfterClosingBrace = h.Trimmed == "}" && isControl
}
