package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/generated/server"
)

func (s *Server) GetSnapshots(
	_ context.Context,
	r server.GetSnapshotsRequestObject,
) (server.GetSnapshotsResponseObject, error) {
	start, end := timeRange(r.Params.Start, r.Params.End)
	snapshots, e := s.store.SnapshotsByTimeRange(
		start,
		end,
		limitOr(r.Params.Limit, 1000),
	)

	if e != nil {
		return server.GetSnapshots500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.GetSnapshots200JSONResponse(toSnapshots(snapshots)), nil
}
