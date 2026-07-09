package service_tester

import (
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/mock_notifier"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	n := mock_notifier.New()

	return &Tester{
		Service: service.New(
			store.New(lite.NewMemory()),
			n,
		),
		Notifier: n,
	}
}
