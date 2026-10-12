package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) SetVirtualLabel(
	_ context.Context,
	r server.SetVirtualLabelRequestObject,
) (server.SetVirtualLabelResponseObject, error) {
	result, e := s.service.SetVirtualLabel(r.Name, r.Key, r.Body.Value)

	if e != nil {
		return server.SetVirtualLabel500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.SetVirtualLabel200JSONResponse(*convert.Label(result)), nil
}
