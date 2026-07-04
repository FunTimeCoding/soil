package service

import (
	"github.com/funtimecoding/go-library/pkg/argocd"
	"github.com/funtimecoding/go-library/pkg/event/notifier"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/nextcloud/usage"
	"github.com/funtimecoding/go-library/pkg/prometheus"
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/board"
	"sync"
)

type Service struct {
	board      *board.Board
	prometheus *prometheus.Client
	usage      *usage.Client
	argocd     *argocd.Client
	notifier   *notifier.Notifier
	logger     *logger.Logger
	mutex      sync.RWMutex
	values     map[string][]string
}
