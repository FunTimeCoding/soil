package runner

import "github.com/funtimecoding/soil/pkg/provision/types/command_run"

func (r *Runner) HostCommand(
	command string,
	host string,
	source string,
) (*command_run.Run, error) {
	return r.provision.HostCommand(command, host, source)
}
