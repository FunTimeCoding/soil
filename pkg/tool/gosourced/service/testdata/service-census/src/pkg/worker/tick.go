package worker

import "example/pkg/service"

func Tick(s *service.Service) error {
	s.Place("bravo")

	return s.Settle()
}
