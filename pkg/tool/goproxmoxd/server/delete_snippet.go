package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) DeleteSnippet(
	_ context.Context,
	r server.DeleteSnippetRequestObject,
) (server.DeleteSnippetResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.DeleteSnippet404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.DeleteSnippet400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.DeleteSnippet500JSONResponse(*s.captureFail(e, e.Error())), nil
	}

	c, e := s.service.SSHClient(instance)

	if e != nil {
		return server.DeleteSnippet500JSONResponse(*s.captureDetail(e)), nil
	}

	c.RemoveFile(snippetPath(r.Name))

	return server.DeleteSnippet200JSONResponse{Status: "deleted"}, nil
}
