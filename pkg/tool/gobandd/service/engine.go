package service

import (
	"github.com/funtimecoding/soil/pkg/band"
	bandConstant "github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/face"
)

func (s *Service) Engine(instance string) (face.Engine, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if c, okay := s.engines[instance]; okay {
		return c, nil
	}

	i, okay := s.Instance(instance)

	if !okay {
		return nil, not_found.New(constant.InstanceKind, instance)
	}

	user := i.User

	if user == "" {
		user = bandConstant.DefaultUser
	}

	result := band.New(i.Locator(), user, i.Password)
	s.engines[instance] = result

	return result, nil
}
