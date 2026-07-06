package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
)

func (s *Server) PostImpressions(
	_ context.Context,
	r server.PostImpressionsRequestObject,
) (server.PostImpressionsResponseObject, error) {
	source := ""

	if r.Body.Source != nil {
		source = *r.Body.Source
	}

	identifier, e := s.service.CreateImpression(r.Body.Content, source)

	if e != nil {
		return server.PostImpressions500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostImpressions200JSONResponse{
		Identifier: int(identifier),
	}, nil
}
