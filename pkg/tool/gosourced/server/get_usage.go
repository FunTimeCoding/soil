package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/server"
)

func (s *Server) GetUsage(
	x context.Context,
	r server.GetUsageRequestObject,
) (server.GetUsageResponseObject, error) {
	m, okay := s.module.Module(r.Params.Module)

	if !okay {
		return server.GetUsage404JSONResponse{
			Error: not_found.New(constant.ModuleKind, r.Params.Module).Error(),
		}, nil
	}

	u, e := s.usage.Usage(m.Directory, r.Params.Package, r.Params.Type)

	if validation.Is(e) {
		return server.GetUsage400JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetUsage500JSONResponse(*s.answer(x, e)), nil
	}

	return server.GetUsage200JSONResponse(usageResponse(u)), nil
}
