package stray_comment

import "strings"

func firstLine(text string) string {
	if index := strings.IndexByte(text, '\n'); index != -1 {
		text = text[:index]
	}

	return strings.TrimSpace(text)
}
