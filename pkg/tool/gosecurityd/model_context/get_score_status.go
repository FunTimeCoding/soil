package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetScoreStatus(
	context.Context,
	mcp.CallToolRequest,
	struct{},
) (*mcp.CallToolResult, error) {
	return response.SuccessAny(convert.ScoreStatus(s.score.Status()))
}
