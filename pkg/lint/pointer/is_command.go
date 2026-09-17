package pointer

import "strings"

func IsCommand(s string) bool {
	if len(s) < 2 || s[0] != '/' {
		return false
	}

	name := s[1:]

	if name[0] < 'a' || name[0] > 'z' {
		return false
	}

	return !strings.ContainsFunc(
		name,
		func(r rune) bool {
			switch {
			case r >= 'a' && r <= 'z':
				return false
			case r >= '0' && r <= '9':
				return false
			case r == '-' || r == '_' || r == ':':
				return false
			}

			return true
		},
	)
}
