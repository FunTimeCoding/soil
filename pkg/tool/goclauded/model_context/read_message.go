package model_context

import (
	"context"
	"fmt"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) readMessage(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ReadMessage,
) (*mcp.CallToolResult, error) {
	if _, e := s.resolveCaller(x, constant.ReadMessage); e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	if len(a.Identifiers) == 0 {
		return response.Fail("identifiers is required")
	}

	var identifiers []uint

	for _, v := range a.Identifiers {
		identifiers = append(identifiers, uint(v))
	}

	found, e := s.service.ReadMessages(identifiers)

	if e != nil {
		return s.captureFail(e, library.UnexpectedError)
	}

	var lines []string

	for _, m := range found {
		lines = append(lines, messageHeader(m), m.Body, "")
	}

	if missing := message.Missing(identifiers, found); len(missing) > 0 {
		lines = append(
			lines,
			fmt.Sprintf(constant.MessageNotFound, join.Comma(numbers(missing))),
		)
	}

	return response.Success(join.NewLine(lines))
}
