package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/mock_notifier"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/store_tester"
	"testing"
	"time"
)

func New(t *testing.T) *Tester {
	t.Helper()
	n := mock_notifier.New()
	now := time.Now().UTC()
	v := service.New(
		store_tester.NewOption(t, func() time.Time { return now }, n),
		logger.New(t.Context()),
	)
	t.Cleanup(v.Close)

	return &Tester{t: t, Service: v, Notifier: n, now: &now}
}
