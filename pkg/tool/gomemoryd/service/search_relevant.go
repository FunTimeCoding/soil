package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face/search_option"
)

func (s *Service) SearchRelevant(
	query string,
	limit int,
	exclude []string,
) ([]store.SearchResult, error) {
	o := search_option.New(query, "memories", limit)
	o.Exclude = exclude
	results, e := s.searcher.Search(o)

	if e != nil {
		return nil, e
	}

	var matches []store.SearchResult

	for _, r := range results {
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
			store.SearchResult{
				Identifier:  m.Identifier,
				Name:        m.Name,
				Content:     m.Content,
				Description: m.Description,
				Type:        m.Type,
				UpdatedAt:   m.UpdatedAt,
				Rank:        r.Score,
				Tags:        m.Tags,
			},
		)
	}

	return matches, nil
}
