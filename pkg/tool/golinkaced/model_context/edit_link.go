package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/types/edit_link_option"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) EditLink(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.EditLink,
) (*mcp.CallToolResult, error) {
	result, e := s.service.EditLink(
		a.Identifier,
		edit_link_option.Option{
			Name:        a.Name,
			Link:        a.Link,
			Description: a.Description,
			Tags:        a.Tags,
			AddTags:     a.AddTags,
			RemoveTags:  a.RemoveTags,
			Lists:       a.Lists,
			AddLists:    a.AddLists,
			RemoveLists: a.RemoveLists,
		},
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(
		map[string]any{
			"identifier": result.Identifier,
			"name":       result.Title,
			"link":       result.Link,
		},
	)
}
