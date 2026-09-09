package service

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/face"
	"sync"
)

type Service struct {
	board      *board.Board
	prometheus face.MetricSource
	usage      face.UsageSource
	argocd     face.ApplicationSource
	notifier   *notifier.Notifier
	logger     *logger.Logger
	mutex      sync.RWMutex
	values     map[string][]string
}
