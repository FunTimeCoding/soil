package service

import (
	"github.com/funtimecoding/soil/pkg/time/window"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/snapshot"
)

func (s *Service) Snapshots(
	w *window.Window,
	limit int,
) ([]snapshot.Snapshot, error) {
	return s.store.Snapshots(w.Bound(s.clock(), constant.QuerySpan), limit)
}
