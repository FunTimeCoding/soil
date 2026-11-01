package markup_tester

func nameEnd(
	s string,
	start int,
) int {
	i := start

	for i < len(s) {
		c := s[i]

		if c == '-' || c == '_' || c >= '0' && c <= '9' ||
			c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			i++

			continue
		}

		break
	}

	return i
}
