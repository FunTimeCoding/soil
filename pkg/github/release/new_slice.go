package release

import "github.com/google/go-github/v90/github"

func NewSlice(v []*github.RepositoryRelease) []*Release {
	var result []*Release

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
