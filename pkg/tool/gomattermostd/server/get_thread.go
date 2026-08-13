package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/server"
)

func (s *Server) GetThread(
	_ context.Context,
	r server.GetThreadRequestObject,
) (server.GetThreadResponseObject, error) {
	parent, e := s.client.FindPost(r.Identifier)

	if e != nil {
		return server.GetThread404JSONResponse(*clientError("post not found")), nil
	}

	replies, f := s.client.Thread(parent)

	if f != nil {
		return server.GetThread500JSONResponse(*s.captureDetail(f)), nil
	}

	return server.GetThread200JSONResponse(convertPosts(replies)), nil
}
