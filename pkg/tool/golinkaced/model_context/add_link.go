package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) AddLink(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.AddLink,
) (*mcp.CallToolResult, error) {
	listIdentifier, e := s.resolveList(a.List)

	if e != nil {
		return s.captureDetail(e)
	}

	name := a.Name

	if name == "" {
		name = a.Link
	}

	result, f := s.client.CreateLink(a.Link, name, listIdentifier, a.Tags)

	if f != nil {
		return s.captureDetail(f)
	}

	return response.SuccessAny(
		map[string]any{
			"identifier": result.Identifier,
			"name":       result.Title,
			"link":       result.Link,
		},
	)
}
