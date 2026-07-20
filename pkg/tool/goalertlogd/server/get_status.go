package server

import (
	"context"
	libraryConstant "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/server"
)

func (s *Server) GetStatus(
	_ context.Context,
	_ server.GetStatusRequestObject,
) (server.GetStatusResponseObject, error) {
	count, e := s.store.Count()

	if e != nil {
		return server.GetStatus500JSONResponse(
			*s.captureFail(e, libraryConstant.UnexpectedError),
		), nil
	}

	result := server.GetStatus200JSONResponse{TotalRecords: count}
	lastPoll := s.worker.LastPoll()

	if !lastPoll.IsZero() {
		result.LastPoll = new(lastPoll.Format(constant.DateFormat))
	}

	return result, nil
}
