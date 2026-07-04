package service

import (
	"github.com/funtimecoding/go-library/pkg/argocd"
	"github.com/funtimecoding/go-library/pkg/event/notifier"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/nextcloud/usage"
	"github.com/funtimecoding/go-library/pkg/prometheus"
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/board"
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
