package runner

import (
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/types/command_run"
)

func (r *Runner) HostCommand(
	command string,
	host string,
	source string,
) (*command_run.Run, error) {
	if r.hostCommandFunction == nil {
		return nil, not_configured.Format(constant.HostCommandMissing)
	}

	r.hostCommand.Lock()
	defer r.hostCommand.Unlock()

	if r.stopped() {
		return nil, unavailable.Format(constant.RunnerStopped)
	}

	return r.hostCommandFunction(command, host, source)
}
