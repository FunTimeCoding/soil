package spacing

import "strings"

func (s *Spacing) trackBlocks(h *shape) {
	if h.elseContinuation {
		if len(s.blockStack) > 0 {
			s.blockStack = s.blockStack[:len(s.blockStack)-1]
		}

		if h.endsWithBrace {
			s.blockStack = append(s.blockStack, true)
		}

		s.pendingControl = false
	} else if strings.HasPrefix(h.trimmed, "}") && h.endsWithBrace {
		control := s.pendingControl

		if len(s.blockStack) > 0 {
			control = control || s.blockStack[len(s.blockStack)-1]
			s.blockStack = s.blockStack[:len(s.blockStack)-1]
		}

		s.blockStack = append(s.blockStack, control)
		s.pendingControl = false
	} else if h.endsWithBrace {
		s.blockStack = append(s.blockStack, h.controlStart || s.pendingControl)
		s.pendingControl = false
	} else if h.controlStart {
		s.pendingControl = true
	} else if !h.blank && !h.closingBrace && s.parenDepth == 0 {
		pastContinues := strings.HasSuffix(h.pastTrimmed, "&&") ||
			strings.HasSuffix(h.pastTrimmed, "||")

		if !s.pendingControl || !pastContinues {
			s.pendingControl = false
		}
	}
}
