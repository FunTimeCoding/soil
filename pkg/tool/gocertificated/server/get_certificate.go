package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
)

func (s *Server) GetCertificate(
	_ context.Context,
	r server.GetCertificateRequestObject,
) (server.GetCertificateResponseObject, error) {
	result, e := s.service.BySerial(r.Serial)

	if not_found.Is(e) {
		return server.GetCertificate404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetCertificate500JSONResponse(*s.answer(e)), nil
	}

	return server.GetCertificate200JSONResponse(
		*convert.Certificate(result),
	), nil
}
