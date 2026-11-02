package dictionary

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/text/constant"
	"slices"
	"strings"
)

func Stems(word string) []string {
	var result []string

	for suffix, endings := range constant.Inflections {
		if !strings.HasSuffix(word, suffix) {
			continue
		}

		stem := strings.TrimSuffix(word, suffix)

		for _, e := range endings {
			if b := join.Empty(stem, e); len(b) >= constant.MinimumStem {
				result = append(result, b)
			}
		}

		if n := len(stem); constant.Doubling[suffix] &&
			n > constant.MinimumStem &&
			stem[n-1] == stem[n-2] {
			result = append(result, stem[:n-1])
		}
	}

	slices.Sort(result)

	return slices.Compact(result)
}
