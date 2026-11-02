package document

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
)

func New(
	s *store.Store,
	q face.Client,
	g *peer.Guard,
	l *logger.Logger,
) *Service {
	return &Service{store: s, query: q, queryGuard: g, logger: l}
}
