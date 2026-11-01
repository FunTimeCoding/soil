package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (s *Service) ListCompletions() ([]client.SearchResult, error) {
	limit := 20
	var outcome *client.ListOutcome

	if e := s.guard.Attempt(
		context.Background(),
		func(x context.Context) error {
			var f error
			outcome, f = s.query.GetList(
				x,
				&client.GetListParams{
					Collection: "completions",
					Metadata: &map[string]string{
						"source_type": "session-completion",
					},
					Limit: &limit,
					Full:  new(true),
				},
			)

			return f
		},
	); e != nil {
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
