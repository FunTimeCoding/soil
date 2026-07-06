package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/habitica/request"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createTask(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	taskType, f := r.RequireString(constant.TaskType)

	if f != nil {
		return response.Fail("task_type is required: %v", f)
	}

	text, g := r.RequireString(constant.Text)

	if g != nil {
		return response.Fail("text is required: %v", g)
	}

	body := request.New()
	body.Type = taskType
	body.Text = text
	body.Notes = r.GetString(constant.Notes, "")
	result, h := s.habitica.CreateTask(body)

	if h != nil {
		return s.captureDetail(h)
	}

	return response.SuccessAny(convert.Task(result))
}
