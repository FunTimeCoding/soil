package service

import (
	"github.com/funtimecoding/soil/pkg/argocd"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/nextcloud/usage"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
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
