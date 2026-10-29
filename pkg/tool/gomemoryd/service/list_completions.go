package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (s *Service) ListCompletions() ([]client.SearchResult, error) {
	limit := 20
	outcome, e := s.query.GetList(
		context.Background(),
		&client.GetListParams{
			Collection: "completions",
			Metadata:   &map[string]string{"source_type": "session-completion"},
			Limit:      &limit,
			Full:       new(true),
		},
	)

	if e != nil {
		return nil, e
	}

	seen := map[string]bool{}
	var filtered []client.SearchResult

	for _, r := range outcome.Results {
		slug := sessionSlug(r.Path)
		seen[slug] = true

		if len(seen) > 5 {
			break
		}

		filtered = append(filtered, r)
	}

	return filtered, nil
}
