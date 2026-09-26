package semver

import (
	"golang.org/x/mod/semver"
	"sort"
)

func SortDescending(tags []string) {
	sort.Slice(
		tags,
		func(i, j int) bool {
			return semver.Compare(
				Prefix(tags[i]),
				Prefix(tags[j]),
			) > 0
		},
	)
}
