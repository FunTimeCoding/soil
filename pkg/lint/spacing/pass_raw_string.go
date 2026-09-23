package spacing

import "strings"

func (s *Spacing) passRawString(line string) bool {
	was := s.inBacktick

	if strings.Count(line, "`")%2 == 1 {
		s.inBacktick = !s.inBacktick
	}

	if !was {
		return false
	}

	s.report.ChangedLine(line)
	s.pastLine = line
	s.pastWasBlank = strings.TrimSpace(line) == ""

	return true
}
