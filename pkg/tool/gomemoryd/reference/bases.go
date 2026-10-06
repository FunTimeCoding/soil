package reference

import (
	"github.com/funtimecoding/soil/pkg/strings/split"
	"strings"
)

func Bases(value string) []string {
	var result []string

	for _, entry := range split.Comma(value) {
		if entry = strings.TrimSpace(entry); entry != "" {
			result = append(result, entry)
		}
	}

	return result
}
