package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
)

func (s *Server) GetDocument(
	_ context.Context,
	r server.GetDocumentRequestObject,
) (server.GetDocumentResponseObject, error) {
	document, _, e := s.service.GetDocument(r.Params.Path)

	if validation.Is(e) {
		return server.GetDocument400JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.GetDocument404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetDocument500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	result := server.GetDocument200JSONResponse{
		Collection:  document.Collection,
		FilePath:    document.FilePath,
		Hash:        document.Hash,
		Path:        document.Path,
		Title:       document.Title,
		VirtualPath: document.VirtualPath,
	}

	if document.Body != "" {
		result.Body = new(string(document.Body))
	}

	if document.Context != "" {
		result.Context = new(string(document.Context))
	}

	return result, nil
}
