package service

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/face"
)

func New(
	b *board.Board,
	p face.MetricSource,
	u face.UsageSource,
	a face.ApplicationSource,
	n *notifier.Notifier,
	l *logger.Logger,
) *Service {
	return &Service{
		board:      b,
		prometheus: p,
		usage:      u,
		argocd:     a,
		notifier:   n,
		logger:     l,
		values:     map[string][]string{},
	}
}
