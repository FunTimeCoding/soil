package spacing

func (s *Spacing) trackParentheses(line string) {
	for _, c := range line {
		switch c {
		case '(':
			s.parenDepth++
		case ')':
			if s.parenDepth > 0 {
				s.parenDepth--
			}
		}
	}
}
