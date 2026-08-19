package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
)

func (s *Server) PostSigningRequest(
	_ context.Context,
	r server.PostSigningRequestRequestObject,
) (server.PostSigningRequestResponseObject, error) {
	result, e := s.service.SignRequest(r.Body)

	if e != nil {
		return server.PostSigningRequest500JSONResponse(
			*s.captureFail(e, constant.SignFail),
		), nil
	}

	return server.PostSigningRequest200JSONResponse(
		*convert.Certificate(result),
	), nil
}
