package image

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"golang.org/x/mod/semver"
)

func Latest(v []*Image) *Image {
	result := v[0]

	for _, e := range v {
		current := e.Version()
		// skip latest
		if current == constant.LatestVersion {
			continue
		}

		if semver.Compare(current, result.Version()) > 0 {
			result = e
		}
	}

	return result
}
