package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Search(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Search,
) (*mcp.CallToolResult, error) {
	r, e := s.service.Search(a.Query, a.Type)

	if e != nil {
		return s.captureDetail(e)
	}

	result := map[string]any{}

	if r.Links != nil {
		links := make([]map[string]any, len(r.Links))

		for i, l := range r.Links {
			links[i] = map[string]any{
				"identifier": l.Identifier,
				"name":       l.Title,
				"link":       l.Link,
			}
		}

		result["links"] = links
	}

	if r.Lists != nil {
		lists := make([]map[string]any, len(r.Lists))

		for i, l := range r.Lists {
			lists[i] = map[string]any{
				"identifier": l.Identifier,
				"name":       l.Name,
			}
		}

		result["lists"] = lists
	}

	if r.Tags != nil {
		tags := make([]map[string]any, len(r.Tags))

		for i, t := range r.Tags {
			tags[i] = map[string]any{
				"identifier": t.Identifier,
				"name":       t.Name,
			}
		}

		result["tags"] = tags
	}

	return response.SuccessAny(result)
}
