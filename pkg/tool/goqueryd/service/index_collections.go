package service

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/service/index_result"

func (s *Service) IndexCollections(collection string) ([]*index_result.IndexResult, error) {
	status, e := s.store.Status()

	if e != nil {
		return nil, e
	}

	var results []*index_result.IndexResult

	for _, c := range status.Collections {
		if c.Type == "push" {
			continue
		}

		if collection != "" && c.Name != collection {
			continue
		}

		r := s.store.Index(c.Name)
		results = append(
			results,
			index_result.New(
				c.Name,
				r.Indexed,
				r.Updated,
				r.Unchanged,
				r.Removed,
			),
		)
	}

	return results, nil
}
