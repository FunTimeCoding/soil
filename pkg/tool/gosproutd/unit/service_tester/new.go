package service_tester

import (
	"github.com/funtimecoding/soil/pkg/event/mock_notifier"
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
		store_tester.NewStore(t, func() time.Time { return now }),
		n,
	)

	return &Tester{t: t, Service: v, Notifier: n, now: &now}
}
