package jellyfin

import "strings"

func sameIdentifier(
	a string,
	b string,
) bool {
	return strings.EqualFold(
		strings.ReplaceAll(a, "-", ""),
		strings.ReplaceAll(b, "-", ""),
	)
}
