package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) DeleteSnippet(
	_ context.Context,
	r server.DeleteSnippetRequestObject,
) (server.DeleteSnippetResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.DeleteSnippet400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.SSHClient(instance)

	if e != nil {
		return server.DeleteSnippet500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	c.RemoveFile(constant.SnippetPath(r.Name))

	return server.DeleteSnippet200JSONResponse{
		Status: "deleted",
	}, nil
}
