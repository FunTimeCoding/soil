package text

import "strings"

func TrimToLastSentence(text string) string {
	i := strings.LastIndexAny(text, ".!?")

	if i == -1 || i == len(text)-1 {
		return text
	}

	return text[:i+1]
}
