package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ListStorages(
	_ context.Context,
	r server.ListStoragesRequestObject,
) (server.ListStoragesResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.ListStorages400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.ListStorages500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	storages, e := s.service.ListStorages(c, r.Name)

	if e != nil {
		return server.ListStorages500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	var result server.ListStorages200JSONResponse

	for _, v := range storages {
		result = append(
			result,
			server.Storage{
				Name:    v.Name,
				Type:    &v.Type,
				Content: &v.Content,
				Enabled: new(v.Enabled == 1),
				Shared:  new(v.Shared == 1),
				Active:  new(v.Active == 1),
				Avail:   new(int64(v.Avail)),
				Used:    new(int64(v.Used)),
				Total:   new(int64(v.Total)),
			},
		)
	}

	return result, nil
}
