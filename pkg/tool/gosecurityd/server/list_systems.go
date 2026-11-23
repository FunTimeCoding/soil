package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
)

func (s *Server) ListSystems(
	x context.Context,
	r server.ListSystemsRequestObject,
) (server.ListSystemsResponseObject, error) {
	p := r.Params
	selected, e := convert.Selection(p.Show, p.Label, p.Search, p.Score, 0)

	if e != nil {
		return server.ListSystems400JSONResponse{Error: e.Error()}, nil
	}

	b, f := s.system.Board(selected)

	if validation.Is(f) {
		return server.ListSystems400JSONResponse{Error: f.Error()}, nil
	}

	if f != nil {
		return server.ListSystems500JSONResponse(*s.answer(x, f)), nil
	}

	return server.ListSystems200JSONResponse(convert.SystemSummaries(b.Rows)), nil
}
