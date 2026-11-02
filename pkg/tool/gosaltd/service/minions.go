package service

import (
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
)

func (s *Service) Minions() ([]response.Minion, error) {
	if !s.runner.Connected() {
		return nil, constant.ErrorNotConnected
	}

	return s.salt.Minions()
}
