package state_tree

import "strings"

func targetHeader(line string) (string, bool) {
	result := strings.TrimSpace(line)

	if strings.HasPrefix(result, "#") {
		return "", false
	}

	if i := strings.Index(result, " #"); i >= 0 {
		result = strings.TrimSpace(result[:i])
	}

	if !strings.HasSuffix(result, ":") {
		return "", false
	}

	result = strings.Trim(strings.TrimSuffix(result, ":"), `'"`)

	return result, result != ""
}
