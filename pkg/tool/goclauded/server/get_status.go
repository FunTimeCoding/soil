package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) GetStatus(
	x context.Context,
	_ server.GetStatusRequestObject,
) (server.GetStatusResponseObject, error) {
	findings, e := s.roster.Findings()

	if e != nil {
		return server.GetStatus500JSONResponse(
			*s.captureFail(x, e, soil.UnexpectedError),
		), nil
	}

	return server.GetStatus200JSONResponse{
		Findings: statusFindings(findings),
	}, nil
}
