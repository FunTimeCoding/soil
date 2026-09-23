package spacing

func (s *Spacing) step(
	line string,
	number int,
) {
	if s.passRawString(line) {
		return
	}

	h := s.shapeOf(line, number)
	s.trackBlocks(h)
	s.trackParentheses(line)

	if !h.blank && s.pendingBlank && s.decideHeldBlank(h) {
		return
	}

	s.requireBlanks(h)

	if h.blank && s.pastWasBlank {
		s.extraneousBlank(h)

		return
	}

	if h.blank {
		s.holdBlank(number)

		return
	}

	s.report.ChangedLine(line)
	s.closeBlock(h)
	s.pastLine = line
	s.pastWasBlank = h.blank
}
