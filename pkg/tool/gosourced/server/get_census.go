package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/server"
)

func (s *Server) GetCensus(
	x context.Context,
	r server.GetCensusRequestObject,
) (server.GetCensusResponseObject, error) {
	m, okay := s.module.Module(r.Params.Module)

	if !okay {
		return server.GetCensus404JSONResponse{
			Error: not_found.New(constant.ModuleKind, r.Params.Module).Error(),
		}, nil
	}

	c, e := s.census.Census(m.Directory, r.Params.Package, r.Params.Type)

	if validation.Is(e) {
		return server.GetCensus400JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetCensus500JSONResponse(*s.answer(x, e)), nil
	}

	return server.GetCensus200JSONResponse(censusResponse(c)), nil
}
