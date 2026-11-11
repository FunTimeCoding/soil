package service

import (
	"cmp"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"maps"
	"slices"
	"strings"
	"time"
)

func latestCompletions(
	results []client.SearchResult,
	sessionLimit int,
) []client.SearchResult {
	latest := map[string]client.SearchResult{}
	stamp := map[string]time.Time{}

	for _, r := range results {
		slug := sessionSlug(r.Path)
		at := completedAt(&r)

		if _, found := latest[slug]; found && !at.After(stamp[slug]) {
			continue
		}

		latest[slug] = r
		stamp[slug] = at
	}

	sessions := slices.SortedFunc(
		maps.Keys(latest),
		func(
			a string,
			b string,
		) int {
			return cmp.Or(stamp[b].Compare(stamp[a]), strings.Compare(a, b))
		},
	)
	var selected []client.SearchResult

	for _, slug := range sessions[:min(sessionLimit, len(sessions))] {
		selected = append(selected, latest[slug])
	}

	return selected
}
