package service_tester

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/mock_notifier"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	n := mock_notifier.New()

	return &Tester{
		Service: service.New(
			store.New(
				filepath.Join(t.TempDir(), constant.TestDatabase),
			),
			n,
		),
		Notifier: n,
	}
}
