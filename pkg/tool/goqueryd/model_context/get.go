package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
	"strings"
)

func (s *Server) get(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	path, e := q.RequireString(constant.Path)

	if e != nil {
		return response.Fail("path is required: %v", e)
	}

	document, similar, f := s.service.GetDocument(path)

	if f != nil {
		return s.captureFail(f, "document lookup failed")
	}

	if document != nil {
		return response.Success(notation.MarshalIndent(document))
	}

	if len(similar) == 0 {
		return response.Fail("document not found: %s", path)
	}

	return response.Fail(
		"document not found: %s\n\nDid you mean:\n%s",
		path,
		fmt.Sprintf("  %s", strings.Join(similar, "\n  ")),
	)
}
