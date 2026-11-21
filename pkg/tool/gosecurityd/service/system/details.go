package system

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"

func (s *Service) Details() ([]*system_result.Detail, error) {
	systems, e := s.store.Systems()

	if e != nil {
		return nil, e
	}

	result := make([]*system_result.Detail, 0, len(systems))

	for i := range systems {
		d, f := s.assess(&systems[i])

		if f != nil {
			return nil, f
		}

		result = append(result, d)
	}

	return result, nil
}
