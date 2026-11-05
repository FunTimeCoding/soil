package service

import "github.com/funtimecoding/soil/pkg/opnsense/lease"

func (s *Service) Leases(phrase string) ([]*lease.Lease, error) {
	return guarded(
		s.guard,
		func() ([]*lease.Lease, error) {
			return s.client.Leases(phrase)
		},
	)
}
