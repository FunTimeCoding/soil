package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateChoiceSet(
	_ context.Context,
	r server.CreateChoiceSetRequestObject,
) (server.CreateChoiceSetResponseObject, error) {
	result, e := s.client.CreateChoiceSet(r.Body.Name, r.Body.Choices)

	if e != nil {
		return server.CreateChoiceSet500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateChoiceSet201JSONResponse(*convert.ChoiceSet(result)), nil
}
