package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/server"
)

func (s *Server) GetStatus(
	_ context.Context,
	_ server.GetStatusRequestObject,
) (server.GetStatusResponseObject, error) {
	status, e := s.service.Status()

	if e != nil {
		return server.GetStatus500JSONResponse(*s.answer(e)), nil
	}

	result := server.GetStatus200JSONResponse{
		TotalRecords:     status.Total,
		LatestIdentifier: int(status.LatestIdentifier),
	}
	lastPoll := s.worker.LastPoll()

	if !lastPoll.IsZero() {
		result.LastPoll = new(lastPoll.Format(constant.DateFormat))
	}

	return result, nil
}
