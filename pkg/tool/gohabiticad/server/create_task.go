package server

import (
	"context"
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

	result, e := s.habitica.CreateTask(c)

	if e != nil {
		return server.CreateTask500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.CreateTask201JSONResponse(
		*convert.Task(result),
	), nil
}
