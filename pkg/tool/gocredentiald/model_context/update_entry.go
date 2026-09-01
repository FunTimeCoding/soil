package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) updateEntry(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Update,
) (*mcp.CallToolResult, error) {
	fields, invalid := parseFields(a.Fields)

	if invalid != "" {
		return response.Fail("field is not KEY=VALUE: %s", invalid)
	}

	if e := s.service.Update(a.Identifier, fields); e != nil {
		return s.captureFail(e, constant.UnexpectedError)
	}

	return response.Success("updated")
}
