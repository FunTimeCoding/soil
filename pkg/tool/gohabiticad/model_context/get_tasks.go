package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getTasks(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	taskType := r.GetString(constant.TaskType, "")
	result, e := s.habitica.Tasks(taskType)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(convert.Tasks(result))
}
