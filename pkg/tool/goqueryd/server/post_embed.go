package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
)

func (s *Server) PostEmbed(
	x context.Context,
	_ server.PostEmbedRequestObject,
) (server.PostEmbedResponseObject, error) {
	result, e := s.service.Embed(x)

	if unavailable.Is(e) {
		return server.PostEmbed503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.PostEmbed500JSONResponse(*s.answer(x, e)), nil
	}

	return server.PostEmbed200JSONResponse{
		Documents: result.Documents,
		Chunks:    result.Chunks,
	}, nil
}
