package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/server"
)

func (s *Server) SetVolume(
	_ context.Context,
	r server.SetVolumeRequestObject,
) (server.SetVolumeResponseObject, error) {
	if e := s.client.SetVolume(r.Identifier, r.Body.Level); e != nil {
		return server.SetVolume500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.SetVolume204Response{}, nil
}
