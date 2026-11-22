package service

import "github.com/funtimecoding/soil/pkg/provision/types/command_run"

func (s *Service) HostCommand(
	command string,
	host string,
	source string,
) (*command_run.Run, error) {
	return s.runner.HostCommand(command, host, source)
}
