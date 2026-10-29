package service

import (
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
)

func (s *Service) Minions() ([]response.Minion, error) {
	c := s.runner.SaltClient()

	if c == nil {
		return nil, constant.ErrorNotConnected
	}

	return c.Minions()
}
