package chunk

import "strings"

func trimmed(
	content string,
	l Line,
) string {
	return strings.TrimSpace(content[l.start:l.end])
}
