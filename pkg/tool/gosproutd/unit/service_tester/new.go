package service_tester

import (
	"github.com/funtimecoding/soil/pkg/event/mock_notifier"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/store_tester"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	n := mock_notifier.New()
	s := store_tester.New(t)

	return &Tester{
		t:        t,
		Service:  service.New(s.Store, n),
		Notifier: n,
		store:    s,
	}
}
