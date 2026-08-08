package release

import "github.com/google/go-github/v90/github"

func New(v *github.RepositoryRelease) *Release {
	return &Release{
		Name:   v.GetTagName(),
		Create: v.GetCreatedAt().Time,
		Raw:    v,
	}
}
