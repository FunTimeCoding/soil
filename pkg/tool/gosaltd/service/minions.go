package service

import "github.com/funtimecoding/soil/pkg/provision/salt/basic/response"

func (s *Service) Minions() ([]response.Minion, error) {
	return guarded(s.guard, s.salt.Minions)
}
