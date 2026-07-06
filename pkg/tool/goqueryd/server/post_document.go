package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	queryConstant "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
)

func (s *Server) PostDocument(
	_ context.Context,
	r server.PostDocumentRequestObject,
) (server.PostDocumentResponseObject, error) {
	metadata := map[string]string{}

	if r.Body.Metadata != nil {
		for key, value := range *r.Body.Metadata {
			metadata[key] = value
		}
	}

	if r.Body.SourceType != nil && *r.Body.SourceType != "" {
		metadata[queryConstant.SourceType] = *r.Body.SourceType
	}

	e := s.service.PushDocument(
		r.Body.Collection,
		r.Body.Path,
		r.Body.Body,
		metadata,
	)

	if e != nil {
		return server.PostDocument500JSONResponse(
			*s.captureFail(
				e,
				constant.UnexpectedError,
			),
		), nil
	}

	return server.PostDocument200Response{}, nil
}
