package service

import "github.com/funtimecoding/soil/pkg/system/types/service"

func (s *Service) Services(origin string) ([]*service.Service, error) {
	all, e := s.reader.Services()

	if e != nil {
		return nil, e
	}

	var result []*service.Service

	for _, v := range all {
		if origin != "" && v.Origin != origin {
			continue
		}

		result = append(result, v)
	}

	return result, nil
}
