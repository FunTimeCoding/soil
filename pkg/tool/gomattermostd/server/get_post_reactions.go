package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/server"
)

func (s *Server) GetPostReactions(
	_ context.Context,
	r server.GetPostReactionsRequestObject,
) (server.GetPostReactionsResponseObject, error) {
	parent, e := s.client.FindPost(r.Identifier)

	if e != nil {
		return server.GetPostReactions404JSONResponse(
			*clientError("post not found"),
		), nil
	}

	result := []string{}

	if parent.HasReactions {
		reactions, f := s.client.Reactions(parent)

		if f != nil {
			return server.GetPostReactions500JSONResponse(
				*s.captureDetail(f),
			), nil
		}

		for _, v := range reactions {
			result = append(result, v.EmojiName)
		}
	}

	return server.GetPostReactions200JSONResponse(result), nil
}
