package module

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"maps"
	"regexp"
	"slices"
	"sort"
)

func publicTypeName(
	rendered string,
	aliases map[string]string,
) string {
	keys := slices.Collect(maps.Keys(aliases))
	sort.Slice(
		keys,
		func(
			i int,
			j int,
		) bool {
			return len(keys[i]) > len(keys[j])
		},
	)

	for _, key := range keys {
		pattern := regexp.MustCompile(
			join.Empty(`(^|[^\w./])`, regexp.QuoteMeta(key), `\b`),
		)
		rendered = pattern.ReplaceAllString(
			rendered,
			join.Empty("${1}", aliases[key]),
		)
	}

	return rendered
}
