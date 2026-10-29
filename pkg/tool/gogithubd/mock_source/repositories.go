package mock_source

import (
	"github.com/google/go-github/v93/github"
	"maps"
	"slices"
)

func (s *Source) Repositories(_ string) ([]*github.Repository, error) {
	var result []*github.Repository

	for _, name := range slices.Sorted(maps.Keys(s.runs)) {
		result = append(result, &github.Repository{Name: new(name)})
	}

	return result, nil
}
