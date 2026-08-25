package collector

import (
	"slices"
	"strings"
)

func sortedContent(content string) string {
	if content == "" {
		return ""
	}

	result := strings.Split(content, ",")
	slices.Sort(result)

	return strings.Join(result, ",")
}
