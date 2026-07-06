package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) GetSnippet(
	_ context.Context,
	r server.GetSnippetRequestObject,
) (server.GetSnippetResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.GetSnippet400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.SSHClient(instance)

	if e != nil {
		return server.GetSnippet500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	b := c.ReadFile(constant.SnippetPath(r.Name))

	return server.GetSnippet200JSONResponse{
		Name:    r.Name,
		Content: string(b),
	}, nil
}
