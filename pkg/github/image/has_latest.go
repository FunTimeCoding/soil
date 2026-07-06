package image

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"slices"
)

func HasLatest(v *Image) bool {
	return slices.Contains(v.Tags, constant.LatestVersion)
}
