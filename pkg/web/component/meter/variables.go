package meter

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func variables(pairs ...string) string {
	var result []string

	for i := 0; i+1 < len(pairs); i += 2 {
		if pairs[i+1] != "" {
			result = append(result, join.Empty(pairs[i], ": ", pairs[i+1]))
		}
	}

	return strings.Join(result, "; ")
}
