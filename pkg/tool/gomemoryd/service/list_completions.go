package service

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"

func (s *Service) ListCompletions() ([]face.SearchResult, error) {
	outcome, e := s.lister.List(
		"completions",
		map[string]string{"source_type": "session-completion"},
		20,
		0,
		true,
	)

	if e != nil {
		return nil, e
	}

	seen := map[string]bool{}
	var filtered []face.SearchResult

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
