package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
)

func (s *Server) PostCertificate(
	_ context.Context,
	r server.PostCertificateRequestObject,
) (server.PostCertificateResponseObject, error) {
	result, key, e := s.service.IssueCertificate(r.Body)

	if not_found.Is(e) {
		return server.PostCertificate404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.PostCertificate500JSONResponse(*s.answer(e)), nil
	}

	response := convert.Certificate(result)
	response.PrivateKey = &key

	return server.PostCertificate200JSONResponse(*response), nil
}
