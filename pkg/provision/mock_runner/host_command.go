package mock_runner

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/generated/server"
	"github.com/funtimecoding/soil/pkg/provision/model/run"
	"github.com/funtimecoding/soil/pkg/provision/types/command_run"
	"github.com/funtimecoding/soil/pkg/provision/types/host_result"
)

func (r *Runner) HostCommand(
	command string,
	host string,
	source string,
) (*command_run.Run, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.fail != nil {
		return nil, r.fail
	}

	request := server.HostCommandRequest{Command: command, Source: source}

	if host != "" {
		request.Host = &host
	}

	r.hostCommands = append(r.hostCommands, request)
	record := run.Stub()
	record.Identifier = uint(len(r.hostCommands))
	record.Scope = command
	record.TriggerSource = source
	record.Status = constant.StoreStatusSuccess

	return command_run.New(
		record,
		[]*host_result.Result{
			host_result.New(host, constant.StoreStatusSuccess, 0, "", ""),
		},
	), nil
}
