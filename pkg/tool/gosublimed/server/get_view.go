package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/server"
)

func (s *Server) GetView(
	_ context.Context,
	r server.GetViewRequestObject,
) (server.GetViewResponseObject, error) {
	v, e := s.client.View(r.Id)
	errors.PanicOnError(e)

	return server.GetView200JSONResponse(*convertView(v)), nil
}
