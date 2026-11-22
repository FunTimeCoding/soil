package server

import (
	"github.com/funtimecoding/soil/pkg/provision/generated/server"
	"github.com/funtimecoding/soil/pkg/provision/types/command_run"
)

func hostCommandRun(r *command_run.Run) server.HostCommandRun {
	hosts := []server.HostCommandResult{}

	for _, h := range r.Hosts {
		hosts = append(
			hosts,
			server.HostCommandResult{
				Host:       h.Host,
				Status:     h.Status,
				ReturnCode: h.ReturnCode,
				Stdout:     h.Stdout,
				Stderr:     h.Stderr,
			},
		)
	}

	return server.HostCommandRun{
		Identifier:  int(r.Record.Identifier),
		Status:      r.Record.Status,
		Output:      r.Record.Output,
		ErrorOutput: r.Record.ErrorOutput,
		Hosts:       hosts,
	}
}
