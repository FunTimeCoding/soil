package label_value

import (
	"github.com/funtimecoding/soil/pkg/strings/split"
	"slices"
	"strings"
)

func Parse(value string) []string {
	var result []string

	for _, v := range split.Comma(value) {
		v = strings.TrimSpace(v)

		if v == "" || slices.Contains(result, v) {
			continue
		}

		result = append(result, v)
	}

	return result
}
