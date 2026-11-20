package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
)

func (s *Server) GetSystemFindings(
	x context.Context,
	r server.GetSystemFindingsRequestObject,
) (server.GetSystemFindingsResponseObject, error) {
	result, e := s.system.Detail(r.Name)

	if not_found.Is(e) {
		return server.GetSystemFindings404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetSystemFindings500JSONResponse(*s.answer(x, e)), nil
	}

	return server.GetSystemFindings200JSONResponse(
		*convert.SystemDetail(result),
	), nil
}
