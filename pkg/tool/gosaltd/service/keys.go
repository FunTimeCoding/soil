package service

import (
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
)

func (s *Service) Keys() (*response.KeysReturn, error) {
	c := s.runner.SaltClient()

	if c == nil {
		return nil, constant.ErrorNotConnected
	}

	return c.Keys()
}
