package container

import "github.com/google/go-github/v90/github"

func New(v *github.Package) *Container {
	var repository string

	if v.Repository != nil {
		repository = v.Repository.GetName()
	}

	return &Container{
		Name:       v.GetName(),
		Repository: repository,
		Raw:        v,
	}
}
