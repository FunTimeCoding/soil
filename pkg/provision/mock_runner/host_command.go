package mock_runner

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/generated/server"
	"github.com/funtimecoding/soil/pkg/provision/model/run"
)

func (r *Runner) HostCommand(
	command string,
	host string,
	source string,
) (*run.Run, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.fail != nil {
		return nil, r.fail
	}

	r.hostCommands = append(
		r.hostCommands,
		server.HostCommandRequest{Command: command, Host: host, Source: source},
	)
	result := run.Stub()
	result.Identifier = uint(len(r.hostCommands))
	result.Scope = command
	result.TriggerSource = source
	result.Status = constant.StoreStatusSuccess

	return result, nil
}
