package server

import (
	"context"
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

	result, e := s.habitica.Tasks(taskType)

	if e != nil {
		return server.GetTasks500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.GetTasks200JSONResponse(convert.Tasks(result)), nil
}
