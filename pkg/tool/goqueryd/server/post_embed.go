package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
)

func (s *Server) PostEmbed(
	_ context.Context,
	_ server.PostEmbedRequestObject,
) (server.PostEmbedResponseObject, error) {
	result, e := s.service.Embed()

	if e != nil {
		return server.PostEmbed500JSONResponse(
			*s.captureFail(
				e,
				constant.UnexpectedError,
			),
		), nil
	}

	return server.PostEmbed200JSONResponse{
		Documents: result.Documents,
		Chunks:    result.Chunks,
	}, nil
}
