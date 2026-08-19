package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
)

func (s *Server) GetPublication(
	_ context.Context,
	_ server.GetPublicationRequestObject,
) (server.GetPublicationResponseObject, error) {
	result, e := s.service.Pending()

	if e != nil {
		return server.GetPublication500JSONResponse(
			*s.captureFail(e, constant.QueryFail),
		), nil
	}

	return server.GetPublication200JSONResponse(
		server.PublicationResponse{Pending: convert.Changes(result)},
	), nil
}
