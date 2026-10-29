package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) GetTags(
	_ context.Context,
	_ server.GetTagsRequestObject,
) (server.GetTagsResponseObject, error) {
	result, e := s.service.Tags()

	if forbidden.Is(e) {
		return server.GetTags403JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetTags503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetTags500JSONResponse(*s.answer(e)), nil
	}

	return server.GetTags200JSONResponse(convert.Tags(result)), nil
}
