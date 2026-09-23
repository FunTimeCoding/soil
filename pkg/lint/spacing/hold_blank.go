package spacing

func (s *Spacing) holdBlank(number int) {
	s.pendingBlank = true
	s.pendingBlankLine = number
	s.pastWasBlank = true
}
