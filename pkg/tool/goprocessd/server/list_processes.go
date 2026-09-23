package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/convert"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/server"
)

func (s *Server) ListProcesses(
	_ context.Context,
	_ server.ListProcessesRequestObject,
) (server.ListProcessesResponseObject, error) {
	return server.ListProcesses200JSONResponse(
		convert.ProcessSlice(s.supervisor.Statuses()),
	), nil
}
