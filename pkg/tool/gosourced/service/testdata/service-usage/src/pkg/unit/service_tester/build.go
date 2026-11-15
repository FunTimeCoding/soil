package service_tester

import "example/pkg/service"

func Build(s *service.Service) int {
	return s.Now()
}
