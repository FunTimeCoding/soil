package system

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"

func (s *Service) details() ([]*system_result.Detail, error) {
	systems, e := s.store.Systems()

	if e != nil {
		return nil, e
	}

	result := make([]*system_result.Detail, 0, len(systems))

	for i := range systems {
		packages, f := s.store.Packages(systems[i].Identifier)

		if f != nil {
			return nil, f
		}

		result = append(result, s.assess(&systems[i], packages))
	}

	return result, nil
}
