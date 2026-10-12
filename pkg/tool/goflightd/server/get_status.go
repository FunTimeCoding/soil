package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/generated/server"
)

func (s *Server) GetStatus(
	_ context.Context,
	_ server.GetStatusRequestObject,
) (server.GetStatusResponseObject, error) {
	events, e := s.store.EventCount()

	if e != nil {
		return server.GetStatus500JSONResponse(
			*s.captureFail(e, soil.UnexpectedError),
		), nil
	}

	snapshots, f := s.store.SnapshotCount()

	if f != nil {
		return server.GetStatus500JSONResponse(
			*s.captureFail(f, soil.UnexpectedError),
		), nil
	}

	marks, g := s.store.MarkCount()

	if g != nil {
		return server.GetStatus500JSONResponse(
			*s.captureFail(g, soil.UnexpectedError),
		), nil
	}

	result := server.GetStatus200JSONResponse{
		Events:    events,
		Snapshots: snapshots,
		Marks:     marks,
	}
	last, h := s.store.LastEventTime()

	if h != nil {
		return server.GetStatus500JSONResponse(
			*s.captureFail(h, soil.UnexpectedError),
		), nil
	}

	if last != nil {
		result.LastEvent = new(last.Format(constant.DateFormat))
	}

	return result, nil
}
