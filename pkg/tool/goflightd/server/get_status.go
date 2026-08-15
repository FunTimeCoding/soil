package server

import (
	"context"
	libraryConstant "github.com/funtimecoding/soil/pkg/constant"
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
			*s.captureFail(e, libraryConstant.UnexpectedError),
		), nil
	}

	snapshots, f := s.store.SnapshotCount()

	if f != nil {
		return server.GetStatus500JSONResponse(
			*s.captureFail(f, libraryConstant.UnexpectedError),
		), nil
	}

	marks, g := s.store.MarkCount()

	if g != nil {
		return server.GetStatus500JSONResponse(
			*s.captureFail(g, libraryConstant.UnexpectedError),
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
			*s.captureFail(h, libraryConstant.UnexpectedError),
		), nil
	}

	if last != nil {
		result.LastEvent = new(last.Format(constant.DateFormat))
	}

	return result, nil
}
