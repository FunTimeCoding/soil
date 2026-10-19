package model_context

import (
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) captureDetail(e error) (*mcp.CallToolResult, error) {
	if f := connection.Classify(e); f != nil {
		return s.captureFail(e, f.Error())
	}

	return response.Answer(s.reporter, e)
}
