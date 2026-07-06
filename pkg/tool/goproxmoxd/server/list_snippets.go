package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ListSnippets(
	_ context.Context,
	r server.ListSnippetsRequestObject,
) (server.ListSnippetsResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.ListSnippets400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.SSHClient(instance)

	if e != nil {
		return server.ListSnippets500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	entries := c.ListDirectory(constant.SnippetDirectory)
	var result server.ListSnippets200JSONResponse

	for _, v := range entries {
		if v.IsDir() {
			continue
		}

		result = append(
			result,
			server.SnippetItem{
				Name: v.Name(),
				Size: new(int64(v.Size())),
			},
		)
	}

	return result, nil
}
