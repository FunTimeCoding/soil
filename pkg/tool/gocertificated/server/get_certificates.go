package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/convert"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
)

func (s *Server) GetCertificates(
	_ context.Context,
	r server.GetCertificatesRequestObject,
) (server.GetCertificatesResponseObject, error) {
	f := store.NewFilter()

	if r.Params.Authority != nil {
		f.Authority = *r.Params.Authority
	}

	if r.Params.Kind != nil {
		f.Kind = string(*r.Params.Kind)
	}

	f.Before = r.Params.ExpiresBefore
	f.Revoked = r.Params.Revoked

	if r.Params.Limit != nil {
		f.Limit = *r.Params.Limit
	}

	result, e := s.store.Certificates(f)

	if e != nil {
		return server.GetCertificates500JSONResponse(
			*s.captureFail(e, constant.QueryFail),
		), nil
	}

	return server.GetCertificates200JSONResponse(
		convert.Certificates(result),
	), nil
}
