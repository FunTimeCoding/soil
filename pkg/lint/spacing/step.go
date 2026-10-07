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

	if !h.Blank && s.pendingBlank && s.decideHeldBlank(h) {
		return
	}

	s.requireBlanks(h)

	if h.Blank && s.pastWasBlank {
		s.extraneousBlank(h)

		return
	}

	if h.Blank {
		s.holdBlank(number)

		return
	}

	s.report.ChangedLine(line)
	s.closeBlock(h)
	s.pastLine = line
	s.pastWasBlank = h.Blank
}
