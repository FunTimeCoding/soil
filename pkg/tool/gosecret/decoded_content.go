package gosecret

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"sort"
	"strings"
)

func DecodedContent(decoded map[string]string) string {
	keys := make([]string, 0, len(decoded))

	for k := range decoded {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	var result string

	for _, k := range keys {
		result = join.Empty(
			result,
			fmt.Sprintf(
				"=== %s ===\n%s\n",
				k,
				strings.TrimRight(decoded[k], "\n"),
			),
		)
	}

	return result
}
