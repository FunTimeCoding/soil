package service

import "github.com/funtimecoding/soil/pkg/provision/salt/basic/response"

func (s *Service) Keys() (*response.KeysReturn, error) {
	return guarded(s.guard, s.salt.Keys)
}
