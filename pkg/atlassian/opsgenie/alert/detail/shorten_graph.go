package detail

import "strings"

func shortenGraph(s string) string {
	if false {
		s = strings.ReplaceAll(s, "&g0.tab=0", "")
	}

	s = strings.ReplaceAll(s, "&g0.explain=0", "")
	s = strings.ReplaceAll(s, "&g0.engine=prometheus", "")
	s = strings.ReplaceAll(s, "&g0.stacked=0", "")
	s = strings.ReplaceAll(s, "&g0.partial_response=0", "")
	s = strings.ReplaceAll(s, "&g0.max_source_resolution=0s", "")
	s = strings.ReplaceAll(s, "&g0.range_input=30m", "")
	s = strings.ReplaceAll(s, "&g0.deduplicate=1", "")
	s = strings.ReplaceAll(s, "&g0.store_matches=%5B%5D", "")
	s = strings.ReplaceAll(s, "&g0.show_exemplars=0", "")

	return s
}
