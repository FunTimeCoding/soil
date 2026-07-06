package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) OpenFile(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.OpenFile,
) (*mcp.CallToolResult, error) {
	if a.FilePath == "" {
		return response.Fail("file_path is required")
	}

	v, e := s.client.OpenFile(a.FilePath)

	if e != nil {
		return s.captureFail(e, constant.UnexpectedError)
	}

	return response.SuccessAny(v)
}
