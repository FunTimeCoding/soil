package model_context

import (
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/linkace/link"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) linkPage(
	links []*link.Link,
	total int,
	limit int,
	offset int,
) (*mcp.CallToolResult, error) {
	result := make([]map[string]any, len(links))

	for i, l := range links {
		result[i] = map[string]any{
			"identifier": l.Identifier,
			"name":       l.Title,
			"link":       l.Link,
		}
	}

	result = paginate(result, limit, offset)

	return response.SuccessAny(map[string]any{"total": total, "links": result})
}
