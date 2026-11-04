package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
)

func (s *Server) GetStatus(
	x context.Context,
	_ server.GetStatusRequestObject,
) (server.GetStatusResponseObject, error) {
	status, e := s.service.Status()

	if e != nil {
		return server.GetStatus500JSONResponse(
			*s.captureFail(x, e, soil.UnexpectedError),
		), nil
	}

	collections := make([]server.CollectionStatus, len(status.Collections))

	for i, c := range status.Collections {
		collections[i] = server.CollectionStatus{
			Name:          c.Name,
			Path:          c.Path,
			Pattern:       c.Pattern,
			DocumentCount: c.DocumentCount,
		}
	}

	return server.GetStatus200JSONResponse{
		Collections:       collections,
		TotalDocuments:    status.TotalDocuments,
		TotalEmbeddings:   status.TotalEmbeddings,
		PendingEmbeddings: status.PendingEmbeddings,
		FailedEmbeddings:  status.FailedEmbeddings,
	}, nil
}
