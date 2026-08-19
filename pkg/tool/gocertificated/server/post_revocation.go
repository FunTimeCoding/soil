package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"time"
)

func (s *Server) PostRevocation(
	_ context.Context,
	r server.PostRevocationRequestObject,
) (server.PostRevocationResponseObject, error) {
	existing, e := s.store.BySerial(r.Serial)

	if e != nil {
		return server.PostRevocation500JSONResponse(
			*s.captureFail(e, constant.QueryFail),
		), nil
	}

	if existing == nil || existing.Revoked != nil {
		return server.PostRevocation404JSONResponse(
			*clientError(constant.CertificateMissing),
		), nil
	}

	if f := s.store.Revoke(r.Serial, time.Now()); f != nil {
		return server.PostRevocation500JSONResponse(
			*s.captureFail(f, constant.RevokeFail),
		), nil
	}

	result, g := s.store.BySerial(r.Serial)

	if g != nil {
		return server.PostRevocation500JSONResponse(
			*s.captureFail(g, constant.QueryFail),
		), nil
	}

	return server.PostRevocation200JSONResponse(
		*convert.Certificate(result),
	), nil
}
