package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/server"
)

func (s *Server) PostCheck(
	x context.Context,
	r server.PostCheckRequestObject,
) (server.PostCheckResponseObject, error) {
	o := option.New()
	o.Root = r.Body.Root
	o.Fix = r.Body.Fix != nil && *r.Body.Fix
	o.Fresh = r.Body.Fresh != nil && *r.Body.Fresh

	if r.Body.Transcript != nil {
		o.Transcript = *r.Body.Transcript
	}

	c, e := s.check.Check(x, o)

	if not_found.Is(e) {
		return server.PostCheck404JSONResponse{Error: e.Error()}, nil
	}

	if conflict.Is(e) {
		return server.PostCheck409JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.PostCheck500JSONResponse(*s.answer(x, e)), nil
	}

	return server.PostCheck200JSONResponse(checkResponse(c)), nil
}
