package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
)

func (s *Server) GetCertificate(
	_ context.Context,
	r server.GetCertificateRequestObject,
) (server.GetCertificateResponseObject, error) {
	result, e := s.store.BySerial(r.Serial)

	if e != nil {
		return server.GetCertificate500JSONResponse(
			*s.captureFail(e, constant.QueryFail),
		), nil
	}

	if result == nil {
		return server.GetCertificate404JSONResponse(
			*clientError(constant.CertificateMissing),
		), nil
	}

	return server.GetCertificate200JSONResponse(
		*convert.Certificate(result),
	), nil
}
