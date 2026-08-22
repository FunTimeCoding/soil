package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/server"
)

func (s *Server) CloseView(
	_ context.Context,
	r server.CloseViewRequestObject,
) (server.CloseViewResponseObject, error) {
	errors.PanicOnError(s.client.CloseView(r.Id))

	return server.CloseView204Response{}, nil
}
