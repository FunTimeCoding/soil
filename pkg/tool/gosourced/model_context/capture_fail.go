package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/load_failure"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) captureFail(
	x context.Context,
	e error,
) (*mcp.CallToolResult, error) {
	if failure, okay := load_failure.As(e); okay {
		text := fmt.Sprintf(
			"module does not load: %s - %v",
			failure.Package,
			failure.Cause,
		)

		if !failure.Settled {
			return response.Fail("%s", text)
		}

		return response.AnswerMessage(x, s.reporter, e, text)
	}

	return response.Answer(x, s.reporter, e)
}
