package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/generated/server"
)

func (s *Server) GetStatus(
	_ context.Context,
	_ server.GetStatusRequestObject,
) (server.GetStatusResponseObject, error) {
	v, e := s.service.Status()

	if e != nil {
		return server.GetStatus500JSONResponse(*s.answer(e)), nil
	}

	result := server.GetStatus200JSONResponse{
		Events:    v.Events,
		Snapshots: v.Snapshots,
		Marks:     v.Marks,
	}

	if v.LastEvent != nil {
		result.LastEvent = new(v.LastEvent.Format(constant.DateFormat))
	}

	return result, nil
}
