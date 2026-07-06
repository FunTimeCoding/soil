package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) GetSessionToolContext(
	_ context.Context,
	r server.GetSessionToolContextRequestObject,
) (server.GetSessionToolContextResponseObject, error) {
	surround := 1

	if r.Params.Surround != nil {
		surround = *r.Params.Surround
	}

	var result []server.ToolContextResult

	for _, c := range s.service.ToolContext(
		r.Identifier,
		r.Params.Filter,
		surround,
	) {
		entry := server.ToolContextResult{
			ToolName: c.ToolName,
		}
		var before []server.SessionMessage

		for _, m := range c.Before {
			before = append(
				before,
				server.SessionMessage{
					Role:      m.Role,
					Text:      m.Text,
					Timestamp: m.Timestamp,
					IsMeta:    new(m.IsMeta),
				},
			)
		}

		if len(before) > 0 {
			entry.Before = &before
		}

		var after []server.SessionMessage

		for _, m := range c.After {
			after = append(
				after,
				server.SessionMessage{
					Role:      m.Role,
					Text:      m.Text,
					Timestamp: m.Timestamp,
					IsMeta:    new(m.IsMeta),
				},
			)
		}

		if len(after) > 0 {
			entry.After = &after
		}

		result = append(result, entry)
	}

	return server.GetSessionToolContext200JSONResponse{Results: result}, nil
}
