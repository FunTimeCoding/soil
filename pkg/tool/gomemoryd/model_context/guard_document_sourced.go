package model_context

import (
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) guardDocumentSourced(
	identifier int64,
) (*mcp.CallToolResult, error) {
	m, e := s.service.GetMemory(identifier)

	if e != nil {
		return nil, nil
	}

	if m.ProvenanceFile == "" {
		return nil, nil
	}

	return response.Fail(
		"memory %d is document-sourced and read-only; edit %s",
		identifier,
		m.ProvenanceFile,
	)
}
