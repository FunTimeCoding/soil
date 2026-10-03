package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/reference"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) checkReferences(
	_ context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	reports, e := s.service.ReferenceCensus()

	if e != nil {
		return s.captureDetail(e)
	}

	if len(reports) == 0 {
		return response.Success(constant.ReferencesClean)
	}

	var lines []string

	for _, r := range reports {
		lines = append(lines, fmt.Sprintf("#%d %s", r.Identifier, r.Name))
		lines = append(lines, reference.Lines(r.Findings)...)
	}

	return response.Success(join.NewLine(lines))
}
