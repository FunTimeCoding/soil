package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) delete(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	collection, e := q.RequireString(constant.Collection)

	if e != nil {
		return response.Fail("collection is required: %v", e)
	}

	path, f := q.RequireString(constant.Path)

	if f != nil {
		return response.Fail("path is required: %v", f)
	}

	deleted, g := s.service.DeleteDocument(collection, path)

	if g != nil {
		return s.captureFail(g, "delete failed")
	}

	if !deleted {
		return response.Fail("not found: %s/%s", collection, path)
	}

	return response.Success("deleted %s/%s", collection, path)
}
