package model_context

import (
	"context"
	"errors"
	"fmt"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) edit(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier, e := q.RequireFloat(constant.Identifier)

	if e != nil {
		return response.Fail("identifier is required: %v", e)
	}

	message, f := q.RequireString(constant.Message)

	if f != nil {
		return response.Fail("message is required: %v", f)
	}

	if identifier < 0 {
		return response.Fail("invalid event identifier")
	}

	_, g := s.resolveCaller(x, constant.EditEvent)

	if g != nil {
		return s.captureFail(g, library.UnexpectedError)
	}

	_, g = s.service.EditEvent(uint(identifier), message)

	if g != nil {
		if errors.Is(g, constant.ErrorEventNotFound) {
			return response.Fail(g.Error())
		}

		return s.captureFail(g, library.UnexpectedError)
	}

	return response.Success(
		fmt.Sprintf("Updated event #%d", int(identifier)),
	)
}
