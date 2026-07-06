package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ListStorageContent(
	_ context.Context,
	r server.ListStorageContentRequestObject,
) (server.ListStorageContentResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.ListStorageContent400JSONResponse(
			*clientError(e),
		), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.ListStorageContent500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	content, e := s.service.ListStorageContent(c, r.Name, r.Storage)

	if e != nil {
		return server.ListStorageContent500JSONResponse(
			*s.captureDetail(e),
		), nil
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
