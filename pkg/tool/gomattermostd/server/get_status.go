package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/stamp"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/server"
)

func (s *Server) GetStatus(
	_ context.Context,
	_ server.GetStatusRequestObject,
) (server.GetStatusResponseObject, error) {
	return server.GetStatus200JSONResponse{
		Name:    constant.Identity.Name(),
		Version: stamp.New().Tag(),
	}, nil
}
