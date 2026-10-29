package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/server"
)

func (s *Server) GetView(
	_ context.Context,
	r server.GetViewRequestObject,
) (server.GetViewResponseObject, error) {
	v, e := s.service.View(r.Id)

	if not_found.Is(e) {
		return server.GetView404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetView500JSONResponse(*s.answer(e)), nil
	}

	return server.GetView200JSONResponse(*convertView(v)), nil
}
