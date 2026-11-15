package markup_tester

func closing(
	s string,
	open int,
) int {
	opener := s[open]
	closer := byte(')')

	if opener == '[' {
		closer = ']'
	}

	depth := 0

	for i := open; i < len(s); i++ {
		switch s[i] {
		case opener:
			depth++
		case closer:
			depth--

			if depth == 0 {
				return i
			}
		}
	}

	return len(s) - 1
}
