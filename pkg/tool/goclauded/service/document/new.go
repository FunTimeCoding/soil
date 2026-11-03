package document

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/metric"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
)

func New(
	s *store.Store,
	q face.Client,
	g *peer.Guard,
	m *metric.Metric,
) *Service {
	return &Service{store: s, query: q, queryGuard: g, metric: m}
}
