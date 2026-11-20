package cross_service_tester

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/store_tester"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	coordinator := base.New(t)
	s := store_tester.New(t)

	return &Tester{
		t:           t,
		Service:     service.New(s.Store, notifier.New()),
		Coordinator: coordinator.Client(),
		Goclauded:   coordinator,
		store:       s,
	}
}
