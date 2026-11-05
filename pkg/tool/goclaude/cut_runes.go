package goclaude

func cutRunes(
	text string,
	limit int,
) (string, bool) {
	runes := []rune(text)

	if len(runes) <= limit {
		return text, false
	}

	return string(runes[:limit]), true
}
