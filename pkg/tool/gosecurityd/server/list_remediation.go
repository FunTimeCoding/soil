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
	p := r.Params
	selected, e := convert.Selection(
		p.Show,
		p.Label,
		p.Search,
		p.Score,
		constant.CriticalScore,
	)

	if e != nil {
		return server.ListRemediation400JSONResponse{Error: e.Error()}, nil
	}

	result, f := s.remediation.List(selected)

	if validation.Is(f) {
		return server.ListRemediation400JSONResponse{Error: f.Error()}, nil
	}

	if f != nil {
		return server.ListRemediation500JSONResponse(*s.answer(x, f)), nil
	}

	return server.ListRemediation200JSONResponse(
		*convert.RemediationList(result),
	), nil
}
