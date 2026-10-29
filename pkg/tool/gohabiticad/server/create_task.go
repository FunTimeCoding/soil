package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/habitica/request"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) CreateTask(
	_ context.Context,
	r server.CreateTaskRequestObject,
) (server.CreateTaskResponseObject, error) {
	c := request.New()
	c.Type = string(r.Body.Type)
	c.Text = r.Body.Text

	if r.Body.Notes != nil {
		c.Notes = *r.Body.Notes
	}

	result, e := s.service.CreateTask(c)

	if forbidden.Is(e) {
		return server.CreateTask403JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.CreateTask503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.CreateTask500JSONResponse(*s.answer(e)), nil
	}

	return server.CreateTask201JSONResponse(*convert.Task(result)), nil
}
