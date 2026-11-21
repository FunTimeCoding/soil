package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
)

func (s *Server) ListFindings(
	x context.Context,
	r server.ListFindingsRequestObject,
) (server.ListFindingsResponseObject, error) {
	result, e := s.system.Groups(convert.Minimum(r.Params.Score))

	if validation.Is(e) {
		return server.ListFindings400JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListFindings500JSONResponse(*s.answer(x, e)), nil
	}

	return server.ListFindings200JSONResponse(convert.FindingGroups(result)), nil
}
