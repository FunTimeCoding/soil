package cross_service_tester

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store"
	"testing"
	"time"
)

func New(t *testing.T) *Tester {
	t.Helper()
	coordinator := base.New(t)
	now := time.Now().UTC()

	return &Tester{
		t: t,
		Service: service.New(
			store.New(lite.NewMemory(), func() time.Time { return now }),
			notifier.New(),
		),
		Coordinator: coordinator.Connector(t),
		Goclauded:   coordinator,
		now:       &now,
	}
}
