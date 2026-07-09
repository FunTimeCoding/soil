package container

import "github.com/google/go-github/v89/github"

func NewSlice(v []*github.Package) []*Container {
	var result []*Container

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
