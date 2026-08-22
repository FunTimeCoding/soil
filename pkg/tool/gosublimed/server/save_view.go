package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/server"
)

func (s *Server) SaveView(
	_ context.Context,
	r server.SaveViewRequestObject,
) (server.SaveViewResponseObject, error) {
	var path string

	if r.Body.FilePath != nil {
		path = *r.Body.FilePath
	}

	errors.PanicOnError(s.client.SaveView(r.Id, path))

	return server.SaveView204Response{}, nil
}
