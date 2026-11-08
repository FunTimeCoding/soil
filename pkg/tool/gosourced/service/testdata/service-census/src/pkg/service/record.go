package service

import "example/pkg/model/receipt"

func (s *Service) record(name string) {
	s.store.Keep(&receipt.Receipt{Name: name})
}
