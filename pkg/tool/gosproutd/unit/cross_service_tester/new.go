package cross_service_tester

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/time/clock_tester"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/store_tester"
	"testing"
	"time"
)

func New(t *testing.T) *Tester {
	t.Helper()
	coordinator := base.New(t)
	c := clock_tester.NewAt(time.Now().UTC())
	v := service.New(store_tester.NewStore(t, c.Now), notifier.New())

	return &Tester{
		t:           t,
		Service:     v,
		Coordinator: coordinator.Client(),
		Goclauded:   coordinator,
		clock:       c,
	}
}
