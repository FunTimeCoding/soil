package document

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/metric"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
)

type Service struct {
	store      *store.Store
	query      face.Client
	queryGuard *peer.Guard
	metric     *metric.Metric
}
