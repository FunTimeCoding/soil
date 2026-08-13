package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/server"
)

func (s *Server) GetChannelPosts(
	_ context.Context,
	r server.GetChannelPostsRequestObject,
) (server.GetChannelPostsResponseObject, error) {
	channel, e := s.client.TeamChannel(r.Name)

	if e != nil {
		return server.GetChannelPosts404JSONResponse(
			*clientError("channel not found"),
		), nil
	}

	posts, f := s.client.PostsSince(channel, r.Params.Since)

	if f != nil {
		return server.GetChannelPosts500JSONResponse(*s.captureDetail(f)), nil
	}

	return server.GetChannelPosts200JSONResponse(convertPosts(posts)), nil
}
