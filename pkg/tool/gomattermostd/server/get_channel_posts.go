package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/server"
)

func (s *Server) GetChannelPosts(
	_ context.Context,
	r server.GetChannelPostsRequestObject,
) (server.GetChannelPostsResponseObject, error) {
	posts, e := s.service.ChannelPosts(r.Name, r.Params.Since)

	if forbidden.Is(e) {
		return server.GetChannelPosts403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.GetChannelPosts404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetChannelPosts503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetChannelPosts500JSONResponse(*s.answer(e)), nil
	}

	return server.GetChannelPosts200JSONResponse(convertPosts(posts)), nil
}
