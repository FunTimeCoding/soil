package style_sheet

func matchingBrace(
	css string,
	open int,
) int {
	depth := 0

	for i := open; i < len(css); i++ {
		switch css[i] {
		case '{':
			depth++
		case '}':
			depth--

			if depth == 0 {
				return i
			}
		}
	}

	return len(css) - 1
}
