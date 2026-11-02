package cross_service_tester

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/store_tester"
	"testing"
	"time"
)

func New(t *testing.T) *Tester {
	t.Helper()
	coordinator := base.New(t)
	now := time.Now().UTC()
	v := service.New(
		store_tester.NewOption(
			t,
			func() time.Time { return now },
			notifier.New(),
		),
		logger.New(t.Context()),
	)
	t.Cleanup(v.Close)

	return &Tester{
		t:           t,
		Service:     v,
		Coordinator: coordinator.Client(t),
		Goclauded:   coordinator,
		now:         &now,
	}
}
