package strings

import (
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"strings"
)

func IndentLines(
	v []string,
	count int,
) []string {
	var result []string
	indent := strings.Repeat(" ", count)

	for _, e := range v {
		result = append(result, key_value.Empty(indent, e))
	}

	return result
}
