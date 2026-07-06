package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/seed"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) reorderSeed(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ReorderSeed,
) (*mcp.CallToolResult, error) {
	seeds := s.service.Seeds()
	var target *seed.Seed

	for _, v := range seeds {
		if v.Name == a.Name {
			found := v
			target = &found

			break
		}
	}

	if target == nil {
		return response.Fail("seed not found: %s", a.Name)
	}

	position := int(a.Position)

	if position < 1 {
		return response.Fail("position must be at least 1")
	}

	if position > len(seeds) {
		position = len(seeds)
	}

	s.service.SetPosition(target.Identifier, position)

	return response.Success(
		fmt.Sprintf("moved %s to position %d", a.Name, position),
	)
}
