package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/server"
)

func (s *Server) OpenFile(
	_ context.Context,
	r server.OpenFileRequestObject,
) (server.OpenFileResponseObject, error) {
	v, e := s.client.OpenFile(r.Body.FilePath)
	errors.PanicOnError(e)

	return server.OpenFile200JSONResponse(*convertView(v)), nil
}
