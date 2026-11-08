package label

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/label"

func (s *Service) LabelsBySessions(
	sessionIdentifiers []string,
) (map[string][]label.Label, error) {
	rows, e := s.store.LabelsBySessions(sessionIdentifiers)

	if e != nil {
		return nil, e
	}

	result := make(map[string][]label.Label, len(rows))

	for identifier, r := range rows {
		result[identifier] = group(r)
	}

	return result, nil
}
