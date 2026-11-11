package pattern

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/match"
	"maps"
	"slices"
)

func resolveQualifiers(
	replacement *match.Pattern,
	querySymbol string,
	anchor string,
	names map[string][]string,
) (map[string]string, error) {
	result, unresolved := qualifierPaths(
		replacement,
		querySymbol,
		anchor,
		names,
	)

	if len(unresolved) == 0 {
		return result, nil
	}

	name := slices.Sorted(maps.Keys(unresolved))[0]

	return nil, fmt.Errorf(
		"cannot resolve package %s (%d candidates) - pass its import line in the replacement source",
		name,
		unresolved[name],
	)
}
