package runner

import "github.com/funtimecoding/soil/pkg/provision/model/run"

func (r *Runner) HostCommand(
	command string,
	host string,
	source string,
) (*run.Run, error) {
	return r.provision.HostCommand(command, host, source)
}
