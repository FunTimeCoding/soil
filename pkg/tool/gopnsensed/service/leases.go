package service

import "github.com/funtimecoding/soil/pkg/opnsense/lease"

func (s *Service) Leases(phrase string) ([]*lease.Lease, error) {
	return s.source.Leases(phrase)
}
