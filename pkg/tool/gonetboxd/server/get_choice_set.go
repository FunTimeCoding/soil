package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) GetChoiceSet(
	_ context.Context,
	r server.GetChoiceSetRequestObject,
) (server.GetChoiceSetResponseObject, error) {
	result, e := s.service.ChoiceSetByName(r.Name)

	if not_found.Is(e) {
		return server.GetChoiceSet404JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.GetChoiceSet403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.GetChoiceSet500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.GetChoiceSet200JSONResponse(*convert.ChoiceSet(result)), nil
}
