package service

import (
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
)

func (s *Service) Keys() (*response.KeysReturn, error) {
	if !s.runner.Connected() {
		return nil, constant.ErrorNotConnected
	}

	return s.salt.Keys()
}
