package state_tree

import "strings"

func listItem(line string) (string, bool) {
	trimmed := strings.TrimSpace(line)

	if !strings.HasPrefix(trimmed, "- ") {
		return "", false
	}

	result := strings.TrimSpace(strings.TrimPrefix(trimmed, "- "))

	if i := strings.Index(result, " #"); i >= 0 {
		result = strings.TrimSpace(result[:i])
	}

	if strings.Contains(result, ": ") {
		return "", false
	}

	result = strings.TrimSuffix(result, ":")
	result = strings.Trim(result, `'"`)

	return result, result != ""
}
