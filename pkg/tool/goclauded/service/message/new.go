package message

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/metric"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stopper"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"time"
)

func New(
	s *store.Store,
	n face.Notifier,
	stop *stopper.Stopper,
	clock func() time.Time,
	m *metric.Metric,
	l *logger.Logger,
) *Service {
	return &Service{
		store:    s,
		notifier: n,
		stop:     stop,
		clock:    clock,
		metric:   m,
		logger:   l,
	}
}
