package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ListSnippets(
	_ context.Context,
	r server.ListSnippetsRequestObject,
) (server.ListSnippetsResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.ListSnippets404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.ListSnippets400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListSnippets500JSONResponse(*s.captureFail(e, e.Error())), nil
	}

	c, e := s.service.SSHClient(instance)

	if e != nil {
		return server.ListSnippets500JSONResponse(*s.captureDetail(e)), nil
	}

	entries := c.ListDirectory(constant.SnippetDirectory)
	var result server.ListSnippets200JSONResponse

	for _, v := range entries {
		if v.IsDir() {
			continue
		}

		result = append(
			result,
			server.SnippetItem{Name: v.Name(), Size: new(int64(v.Size()))},
		)
	}

	return result, nil
}
