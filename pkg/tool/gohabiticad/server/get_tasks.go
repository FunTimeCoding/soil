package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) GetTasks(
	_ context.Context,
	r server.GetTasksRequestObject,
) (server.GetTasksResponseObject, error) {
	var taskType string

	if r.Params.Type != nil {
		taskType = string(*r.Params.Type)
	}

	result, e := s.service.Tasks(taskType)

	if forbidden.Is(e) {
		return server.GetTasks403JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetTasks503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetTasks500JSONResponse(*s.answer(e)), nil
	}

	return server.GetTasks200JSONResponse(convert.Tasks(result)), nil
}
