package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SendText(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SendText,
) (*mcp.CallToolResult, error) {
	if a.Text == "" {
		return response.Fail("text is required")
	}

	e := s.client.SendText(a.SessionIdentifier, a.Text)

	if e != nil {
		return s.captureFail(e, constant.UnexpectedError)
	}

	if a.SendEnter {
		f := s.client.SendKey(a.SessionIdentifier, "enter")

		if f != nil {
			return response.Fail("send enter: %v", f)
		}
	}

	return response.Success("sent: %s", a.Text)
}
