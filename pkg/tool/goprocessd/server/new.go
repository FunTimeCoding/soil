package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/environment"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/process"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/procfile"
)

func New(
	entries []procfile.Entry,
	env *environment.Environment,
	procfilePath string,
	envrcPath string,
	socketPath string,
) *Server {
	maxNameWidth := 0

	for _, entry := range entries {
		if len(entry.Name) > maxNameWidth {
			maxNameWidth = len(entry.Name)
		}
	}

	processes := make([]*process.Process, len(entries))

	for i, entry := range entries {
		processes[i] = process.New(
			entry.Name,
			entry.Command,
			i,
			maxNameWidth,
		)
	}

	return &Server{
		processes:    processes,
		maxNameWidth: maxNameWidth,
		environment:  env,
		procfilePath: procfilePath,
		envrcPath:    envrcPath,
		socketPath:   socketPath,
	}
}
