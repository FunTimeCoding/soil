package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
)

func (s *Server) PostPublication(
	_ context.Context,
	_ server.PostPublicationRequestObject,
) (server.PostPublicationResponseObject, error) {
	commit, change, e := s.service.Publish()

	if e != nil {
		return server.PostPublication500JSONResponse(
			*s.captureFail(e, constant.PublishFail),
		), nil
	}

	result := server.PublicationResponse{Pending: convert.Changes(change)}

	if commit != "" {
		result.Commit = &commit
	}

	return server.PostPublication200JSONResponse(result), nil
}
