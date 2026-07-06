package server

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) CreateSnippet(
	_ context.Context,
	r server.CreateSnippetRequestObject,
) (server.CreateSnippetResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.CreateSnippet400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.SSHClient(instance)

	if e != nil {
		return server.CreateSnippet500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	c.WriteFile(constant.SnippetPath(r.Body.Name), []byte(r.Body.Content))

	return server.CreateSnippet200JSONResponse{
		Volume: fmt.Sprintf("local:snippets/%s", r.Body.Name),
		Status: "created",
	}, nil
}
