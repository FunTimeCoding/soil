package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/server"
)

func (s *Server) GetPostReactions(
	_ context.Context,
	r server.GetPostReactionsRequestObject,
) (server.GetPostReactionsResponseObject, error) {
	result, e := s.service.PostReactions(r.Identifier)

	if forbidden.Is(e) {
		return server.GetPostReactions403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.GetPostReactions404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetPostReactions503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetPostReactions500JSONResponse(*s.answer(e)), nil
	}

	return server.GetPostReactions200JSONResponse(result), nil
}
