package service

import "github.com/funtimecoding/soil/pkg/provision/model/run"

func (s *Service) HostCommand(
	command string,
	host string,
	source string,
) (*run.Run, error) {
	return s.runner.HostCommand(command, host, source)
}
