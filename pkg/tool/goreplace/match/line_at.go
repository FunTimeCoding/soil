package match

import "strings"

func lineAt(
	content string,
	offset int,
) int {
	return strings.Count(content[:offset], "\n") + 1
}
