package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ListStorageContent(
	_ context.Context,
	r server.ListStorageContentRequestObject,
) (server.ListStorageContentResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if not_found.Is(e) {
		return server.ListStorageContent404JSONResponse(*clientError(e)), nil
	}

	if not_selected.Is(e) {
		return server.ListStorageContent400JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListStorageContent500JSONResponse(
			*s.captureFail(e, e.Error()),
		), nil
	}

	content, e := s.service.ListStorageContent(instance, r.Name, r.Storage)

	if e != nil {
		return server.ListStorageContent500JSONResponse(*s.captureDetail(e)), nil
	}

	var result server.ListStorageContent200JSONResponse

	for _, v := range content {
		result = append(
			result,
			server.StorageContentItem{
				Volume: v.Volid,
				Format: &v.Format,
				Size:   new(int64(v.Size)),
			},
		)
	}

	return result, nil
}
