package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	queryConstant "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (s *Service) SearchRelevant(
	query string,
	limit int,
	exclude []string,
) ([]record.SearchResult, error) {
	p := &client.GetSearchParams{
		Query:      query,
		Collection: new(constant.DefaultCollection),
		Limit:      &limit,
		Full:       new(true),
		Mode:       new(client.GetSearchParamsMode(queryConstant.HybridMode)),
	}

	if len(exclude) > 0 {
		p.Exclude = &exclude
	}

	var outcome *client.SearchOutcome

	if e := s.guard.Attempt(
		context.Background(),
		func(x context.Context) error {
			var f error
			outcome, f = s.query.GetSearch(x, p)

			return f
		},
	); e != nil {
		return nil, e
	}
	var matches []record.SearchResult

	for _, r := range outcome.Results {
		identifier, f := extractIdentifier(r.Path)

		if f != nil {
			continue
		}

		m, g := s.store.GetMemory(identifier)

		if g != nil {
			continue
		}

		matches = append(
			matches,
			record.SearchResult{
				Identifier:  m.Identifier,
				Name:        m.Name,
				Content:     m.Content,
				Description: m.Description,
				Type:        m.Type,
				UpdatedAt:   m.UpdatedAt,
				Rank:        float64(r.Score),
				Tags:        m.Tags,
				Metadata:    m.Metadata,
			},
		)
	}

	return matches, nil
}
