package service_tester

import (
	"github.com/funtimecoding/soil/pkg/event/mock_notifier"
	"github.com/funtimecoding/soil/pkg/time/clock_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/store_tester"
	"testing"
	"time"
)

func New(t *testing.T) *Tester {
	t.Helper()
	n := mock_notifier.New()
	c := clock_tester.NewAt(time.Now().UTC())
	v := service.New(store_tester.NewStore(t, c.Now), n)

	return &Tester{t: t, Service: v, Notifier: n, clock: c}
}
