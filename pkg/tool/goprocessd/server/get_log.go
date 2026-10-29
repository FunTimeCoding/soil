package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/server"
)

func (s *Server) GetLog(
	_ context.Context,
	r server.GetLogRequestObject,
) (server.GetLogResponseObject, error) {
	all := r.Params.All != nil && *r.Params.All
	lines, older, e := s.service.Log(r.Name, all)

	if e != nil {
		return server.GetLog404JSONResponse{Error: e.Error()}, nil
	}

	if lines == nil {
		lines = []string{}
	}

	return server.GetLog200JSONResponse{Lines: lines, Older: older}, nil
}
