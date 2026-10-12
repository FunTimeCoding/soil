package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gobandd/face"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/types/status"
)

func (s *Service) Status(instance string) (*status.Status, error) {
	c, e := s.Engine(instance)

	if e != nil {
		return nil, e
	}

	result := status.New(instance)

	for _, read := range []func(face.Engine, *status.Status) error{
		readHost,
		readSetup,
		readScreen,
		readRedirection,
		readConsent,
		readPower,
	} {
		if f := read(c, result); f != nil {
			return nil, f
		}
	}

	return result, nil
}
