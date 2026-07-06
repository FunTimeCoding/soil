package service

import (
	"github.com/funtimecoding/soil/pkg/argocd"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/nextcloud/usage"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
)

func New(
	b *board.Board,
	p *prometheus.Client,
	u *usage.Client,
	a *argocd.Client,
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
