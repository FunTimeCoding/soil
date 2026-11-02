package service

import (
	"github.com/funtimecoding/soil/pkg/errors/dropped"
	"github.com/funtimecoding/soil/pkg/ssh/face"
)

func (s *Service) forgetDropped(
	instance string,
	c face.SSHClient,
	e error,
) error {
	if dropped.Is(e) {
		s.shells.Forget(instance, c)
	}

	return e
}
