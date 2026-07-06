package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/server"
)

func (s *Server) GetStatus(
	_ context.Context,
	_ server.GetStatusRequestObject,
) (server.GetStatusResponseObject, error) {
	count, e := s.store.Count()

	if e != nil {
		return server.GetStatus500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	result := server.GetStatus200JSONResponse{TotalRecords: count}
	lastPoll := s.worker.LastPoll()

	if !lastPoll.IsZero() {
		result.LastPoll = new(lastPoll.Format(DateFormat))
	}

	return result, nil
}
