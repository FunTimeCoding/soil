package service_tester

import (
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/mock_notifier"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store"
	"testing"
	"time"
)

func New(t *testing.T) *Tester {
	t.Helper()
	n := mock_notifier.New()
	now := time.Now().UTC()

	return &Tester{
		t: t,
		Service: service.New(
			store.New(lite.NewMemory(), func() time.Time { return now }),
			n,
		),
		Notifier: n,
		now:      &now,
	}
}
