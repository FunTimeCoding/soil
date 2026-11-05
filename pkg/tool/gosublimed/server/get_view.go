package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/server"
)

func (s *Server) GetView(
	x context.Context,
	r server.GetViewRequestObject,
) (server.GetViewResponseObject, error) {
	v, e := s.service.View(r.Id)

	if not_found.Is(e) {
		return server.GetView404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetView503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetView500JSONResponse(*s.answer(x, e)), nil
	}

	return server.GetView200JSONResponse(*convertView(v)), nil
}
