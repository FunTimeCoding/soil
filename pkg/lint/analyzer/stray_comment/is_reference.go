package stray_comment

import "strings"

func isReference(text string) bool {
	locator, found := strings.CutPrefix(text, "// Reference: ")

	if !found || len(strings.Fields(locator)) != 1 {
		return false
	}

	return strings.HasPrefix(locator, "https://") ||
		strings.HasPrefix(locator, "http://")
}
