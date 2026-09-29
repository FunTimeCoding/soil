package lint

import (
	"github.com/funtimecoding/soil/pkg/strings/split"
	"strings"
)

func declaredValues(entries []string) []string {
	var result []string

	for _, entry := range entries {
		for _, value := range split.Comma(entry) {
			if value = strings.TrimSpace(value); value != "" {
				result = append(result, value)
			}
		}
	}

	return result
}
