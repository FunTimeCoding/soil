package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/server"
)

func (s *Server) EditView(
	_ context.Context,
	r server.EditViewRequestObject,
) (server.EditViewResponseObject, error) {
	all := r.Body.ReplaceAll != nil && *r.Body.ReplaceAll
	v, e := s.client.EditView(r.Id, r.Body.OldString, r.Body.NewString, all)
	errors.PanicOnError(e)

	return server.EditView200JSONResponse(*convertView(v)), nil
}
