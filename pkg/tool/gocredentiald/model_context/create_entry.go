package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createEntry(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Create,
) (*mcp.CallToolResult, error) {
	fields, invalid := parseFields(a.Fields)

	if invalid != "" {
		return response.Fail("field is not KEY=VALUE: %s", invalid)
	}

	identifier, e := s.service.Create(a.Group, a.Title, fields)

	if e != nil {
		return s.captureFail(e, constant.UnexpectedError)
	}

	return response.Success(fmt.Sprintf("created %s", identifier))
}
