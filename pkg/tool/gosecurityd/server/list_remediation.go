package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
)

func (s *Server) ListRemediation(
	x context.Context,
	r server.ListRemediationRequestObject,
) (server.ListRemediationResponseObject, error) {
	m := constant.CriticalScore

	if r.Params.Score != nil {
		m = *r.Params.Score
	}

	result, e := s.remediation.List(m)

	if validation.Is(e) {
		return server.ListRemediation400JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListRemediation500JSONResponse(*s.answer(x, e)), nil
	}

	return server.ListRemediation200JSONResponse(
		*convert.RemediationList(result),
	), nil
}
