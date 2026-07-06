package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) GetSessionMessages(
	_ context.Context,
	r server.GetSessionMessagesRequestObject,
) (server.GetSessionMessagesResponseObject, error) {
	var result []server.SessionMessage

	for _, m := range s.service.Messages(r.Identifier) {
		e := server.SessionMessage{
			Role:      m.Role,
			Text:      m.Text,
			Timestamp: m.Timestamp,
		}

		if m.IsMeta {
			e.IsMeta = &m.IsMeta
		}

		result = append(result, e)
	}

	return server.GetSessionMessages200JSONResponse{Messages: result}, nil
}
