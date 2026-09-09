package image

import (
	"golang.org/x/mod/semver"
	"sort"
)

func Sort(v []*Image) []*Image {
	result := append([]*Image{}, v...)
	sort.SliceStable(
		result,
		func(
			i int,
			j int,
		) bool {
			return semver.Compare(
				result[i].Version(),
				result[j].Version(),
			) > 0
		},
	)

	return result
}
