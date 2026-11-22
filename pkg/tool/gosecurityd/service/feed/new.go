package feed

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/system"
)

func New(
	a face.Client,
	g *peer.Guard,
	y *system.Service,
	l *logger.Logger,
) *Service {
	return &Service{ansible: a, guard: g, system: y, logger: l}
}
