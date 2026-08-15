package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/snapshot"
)

func toSnapshots(v []snapshot.Snapshot) []server.SnapshotResponse {
	result := make([]server.SnapshotResponse, 0, len(v))

	for _, w := range v {
		result = append(
			result,
			server.SnapshotResponse{
				Time:  w.Time.Format(constant.DateFormat),
				Kind:  w.Kind,
				Key:   w.Key,
				Value: w.Value,
			},
		)
	}

	return result
}
