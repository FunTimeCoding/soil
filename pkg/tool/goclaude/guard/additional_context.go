package guard

import "strings"

func additionalContext(
	hint string,
	notes []string,
) string {
	if hint != "" {
		notes = append([]string{hint}, notes...)
	}

	return strings.Join(notes, "\n")
}
