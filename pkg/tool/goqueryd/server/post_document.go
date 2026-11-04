package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/timeout"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
)

func (s *Server) PostDocument(
	x context.Context,
	r server.PostDocumentRequestObject,
) (server.PostDocumentResponseObject, error) {
	metadata := map[string][]string{}

	if r.Body.Metadata != nil {
		for key, value := range *r.Body.Metadata {
			metadata[key] = value
		}
	}

	if r.Body.SourceType != nil && *r.Body.SourceType != "" {
		metadata[constant.SourceType] = []string{*r.Body.SourceType}
	}

	warning, e := s.service.PushDocument(
		x,
		r.Body.Collection,
		r.Body.Path,
		r.Body.Body,
		metadata,
	)

	if timeout.Is(e) {
		return server.PostDocument503JSONResponse(
			*s.captureFail(x, e, e.Error()),
		), nil
	}

	if e != nil {
		return server.PostDocument500JSONResponse(
			*s.captureFail(x, e, soil.UnexpectedError),
		), nil
	}

	result := server.PostDocument200JSONResponse{}

	if warning != "" {
		result.Warning = &warning
	}

	return result, nil
}
