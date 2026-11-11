package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (s *Service) ListCompletions() ([]client.SearchResult, error) {
	collection := "completions"
	limit := 1000
	var outcome *client.ListOutcome

	if e := s.guard.Attempt(
		context.Background(),
		func(x context.Context) error {
			var f error
			outcome, f = s.query.GetList(
				x,
				&client.GetListParams{
					Collection: collection,
					Metadata: &map[string]string{
						"source_type": "session-completion",
					},
					Limit: &limit,
				},
			)

			return f
		},
	); e != nil {
		return nil, e
	}
	var filled []client.SearchResult

	for _, r := range latestCompletions(outcome.Results, 5) {
		var document *client.SearchResult

		if e := s.guard.Attempt(
			context.Background(),
			func(x context.Context) error {
				var f error
				document, f = s.query.GetDocument(
					x,
					&client.GetDocumentParams{
						Path: join.Slash([]string{collection, r.Path}),
					},
				)

				return f
			},
		); e != nil {
			return nil, e
		}
		filled = append(filled, *document)
	}

	return filled, nil
}
