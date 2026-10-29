package service

import "github.com/funtimecoding/soil/pkg/system/types/service"

func (s *Service) Services(origin string) []*service.Service {
	var result []*service.Service

	for _, v := range s.reader.Services() {
		if origin != "" && v.Origin != origin {
			continue
		}

		result = append(result, v)
	}

	return result
}
