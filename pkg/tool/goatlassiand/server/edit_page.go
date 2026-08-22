package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
	"strings"
)

func (s *Server) EditPage(
	_ context.Context,
	r server.EditPageRequestObject,
) (server.EditPageResponseObject, error) {
	title := ""
	message := ""
	draft := false

	if r.Body.Title != nil {
		title = *r.Body.Title
	}

	if r.Body.Message != nil {
		message = *r.Body.Message
	}

	if r.Body.Draft != nil {
		draft = *r.Body.Draft
	}

	result, e := s.service.EditPage(
		r.Identifier,
		r.Body.OldText,
		r.Body.NewText,
		title,
		message,
		draft,
	)

	if e != nil {
		if strings.Contains(e.Error(), constant.OldText) {
			return server.EditPage400JSONResponse(*clientError(e)), nil
		}

		return server.EditPage500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.EditPage200JSONResponse(
		*convert.ConfluencePageDetail(result),
	), nil
}
