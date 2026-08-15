package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/mark"
	"time"
)

func (s *Server) PostMark(
	_ context.Context,
	r server.PostMarkRequestObject,
) (server.PostMarkResponseObject, error) {
	if r.Body == nil || r.Body.Label == "" {
		return server.PostMark400JSONResponse(
			server.Error{Error: "label is required"},
		), nil
	}

	note := ""

	if r.Body.Note != nil {
		note = *r.Body.Note
	}

	v := mark.New(time.Now(), r.Body.Label, note)

	if e := s.store.CreateMark(v); e != nil {
		return server.PostMark500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostMark201JSONResponse(*toMark(*v)), nil
}
