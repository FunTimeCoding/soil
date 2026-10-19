package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) GetSnippet(
	_ context.Context,
	r server.GetSnippetRequestObject,
) (server.GetSnippetResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.GetSnippet404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.GetSnippet400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.GetSnippet500JSONResponse(*s.captureFail(e, e.Error())), nil
	}

	c, e := s.service.SSHClient(instance)

	if e != nil {
		return server.GetSnippet500JSONResponse(*s.captureDetail(e)), nil
	}

	b := c.ReadFile(snippetPath(r.Name))

	return server.GetSnippet200JSONResponse{
		Name:    r.Name,
		Content: string(b),
	}, nil
}
