package model_context

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/mark3labs/mcp-go/mcp"
	"sigs.k8s.io/yaml"
)

func (s *Server) formatMarkup(
	object map[string]any,
	filtered []string,
) (*mcp.CallToolResult, error) {
	b, e := yaml.Marshal(object)

	if e != nil {
		return s.captureFail(e, "marshal yaml")
	}

	if len(filtered) == 0 {
		return response.Success(string(b))
	}

	return response.Success(
		fmt.Sprintf("# filtered: %s\n%s", join.CommaSpace(filtered), b),
	)
}
