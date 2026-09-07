package environment

import "strings"

func harness(base []string) []string {
	var result []string

	for _, entry := range base {
		key, _, found := strings.Cut(entry, "=")

		if found && (key == "PATH" || key == "HOME") {
			result = append(result, entry)
		}
	}

	return result
}
