package spacing

import "strings"

func (s *Spacing) closeBlock(h *shape) {
	if !strings.HasPrefix(h.trimmed, "}") || h.elseContinuation || h.endsWithBrace {
		s.needBlankAfterClosingBrace = false

		return
	}

	isControl := false

	if len(s.blockStack) > 0 {
		isControl = s.blockStack[len(s.blockStack)-1]
		s.blockStack = s.blockStack[:len(s.blockStack)-1]
	}

	s.needBlankAfterClosingBrace = h.trimmed == "}" && isControl
}
