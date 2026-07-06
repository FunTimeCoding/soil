package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"io"
)

func (s *Server) GetPipelineJobOutput(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetPipelineJobOutput,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Job == 0 {
		return response.Fail("job is required")
	}

	r, _, e := s.client.Jobs.GetTraceFile(a.Project, a.Job)

	if e != nil {
		return s.captureDetail(e)
	}

	b, e := io.ReadAll(r)

	if e != nil {
		return s.captureDetail(e)
	}

	return mcp.NewToolResultText(string(b)), nil
}
