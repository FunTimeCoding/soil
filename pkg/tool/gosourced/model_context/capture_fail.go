package model_context

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/load_failure"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) captureFail(
	e error,
	message string,
) (*mcp.CallToolResult, error) {
	if conflict.Is(e) {
		return response.Answer(s.reporter, e)
	}

	if failure, okay := load_failure.As(e); okay {
		message = fmt.Sprintf(
			"module does not load: %s - %v",
			failure.Package,
			failure.Cause,
		)

		if !failure.Settled {
			return response.Fail("%s", message)
		}
	}

	return response.CaptureFail(s.reporter, e, message)
}
