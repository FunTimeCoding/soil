package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) GetSessionsBashDump(
	_ context.Context,
	_ server.GetSessionsBashDumpRequestObject,
) (server.GetSessionsBashDumpResponseObject, error) {
	var result []string

	for _, e := range s.service.Sessions() {
		for _, c := range s.service.ToolCalls(e.Identifier) {
			if c.Name == "Bash" && c.Detail != "" {
				result = append(result, c.Detail)
			}
		}
	}

	return server.GetSessionsBashDump200JSONResponse{Commands: result}, nil
}
