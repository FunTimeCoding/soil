package image

import "github.com/google/go-github/v89/github"

func NewSlice(v []*github.PackageVersion) []*Image {
	var result []*Image

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
