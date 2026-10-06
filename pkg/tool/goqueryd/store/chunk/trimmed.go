package chunk

import "strings"

func trimmed(
	content string,
	l line,
) string {
	return strings.TrimSpace(content[l.start:l.end])
}
