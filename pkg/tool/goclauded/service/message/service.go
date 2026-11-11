package message

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/metric"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stopper"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"time"
)

type Service struct {
	store         *store.Store
	notifier      face.Notifier
	stop          *stopper.Stopper
	clock         func() time.Time
	metric        *metric.Metric
	logger        *logger.Logger
	channelAlways bool
	defaultGrace  time.Duration
}
