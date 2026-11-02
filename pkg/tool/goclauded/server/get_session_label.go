package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) GetSessionLabel(
	_ context.Context,
	r server.GetSessionLabelRequestObject,
) (server.GetSessionLabelResponseObject, error) {
	labels, e := s.label.LabelsBySession(r.Identifier)

	if e != nil {
		return server.GetSessionLabel500JSONResponse(
			*s.captureFail(e, soil.UnexpectedError),
		), nil
	}

	entries := make([]server.LabelEntry, 0, len(labels))

	for _, l := range labels {
		entries = append(entries, server.LabelEntry{Key: l.Key, Value: l.Value})
	}

	return server.GetSessionLabel200JSONResponse(
		server.SessionLabelResponse{Labels: entries},
	), nil
}
