package scan

import "strings"

func (s *Scanner) Match(line string) []string {
	split := splitCase(line)
	lower := strings.ToLower(line)
	lowerSplit := strings.ToLower(split)
	var present [256]bool

	for i := range len(lower) {
		present[lower[i]] = true
	}

	for i := range len(lowerSplit) {
		present[lowerSplit[i]] = true
	}

	var result []string

	for i, term := range s.lower {
		if !present[term[0]] {
			continue
		}

		if !strings.Contains(lower, term) &&
			!strings.Contains(lowerSplit, term) {
			continue
		}

		if s.patterns[i].MatchString(line) ||
			s.patterns[i].MatchString(split) {
			result = append(result, s.terms[i])
		}
	}

	return result
}
