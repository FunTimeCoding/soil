package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
)

func (s *Server) PostCertificate(
	_ context.Context,
	r server.PostCertificateRequestObject,
) (server.PostCertificateResponseObject, error) {
	result, key, e := s.service.IssueCertificate(r.Body)

	if e != nil {
		return server.PostCertificate500JSONResponse(
			*s.captureFail(e, constant.IssueFail),
		), nil
	}

	response := convert.Certificate(result)
	response.PrivateKey = &key

	return server.PostCertificate200JSONResponse(*response), nil
}
