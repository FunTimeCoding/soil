package spacing

import (
	"github.com/funtimecoding/soil/pkg/lint/types/spacing_shape"
	"strings"
)

func (s *Spacing) trackBlocks(h *spacing_shape.Shape) {
	if h.ElseContinuation {
		if len(s.blockStack) > 0 {
			s.blockStack = s.blockStack[:len(s.blockStack)-1]
		}

		if h.EndsWithBrace {
			s.blockStack = append(s.blockStack, true)
		}

		s.pendingControl = false
	} else if strings.HasPrefix(h.Trimmed, "}") && h.EndsWithBrace {
		control := s.pendingControl

		if len(s.blockStack) > 0 {
			control = control || s.blockStack[len(s.blockStack)-1]
			s.blockStack = s.blockStack[:len(s.blockStack)-1]
		}

		s.blockStack = append(s.blockStack, control)
		s.pendingControl = false
	} else if h.EndsWithBrace {
		s.blockStack = append(s.blockStack, h.ControlStart || s.pendingControl)
		s.pendingControl = false
	} else if h.ControlStart {
		s.pendingControl = true
	} else if !h.Blank && !h.ClosingBrace && s.parenDepth == 0 {
		pastContinues := strings.HasSuffix(h.PastTrimmed, "&&") ||
			strings.HasSuffix(h.PastTrimmed, "||")

		if !s.pendingControl || !pastContinues {
			s.pendingControl = false
		}
	}
}
