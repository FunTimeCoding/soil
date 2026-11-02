package image

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"slices"
)

func HasLatest(v *Image) bool {
	return slices.Contains(v.Tags, soil.LatestVersion)
}
