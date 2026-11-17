package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/provision/salt"
	"github.com/funtimecoding/soil/pkg/provision/salt/face"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/option"
	"github.com/prometheus/client_golang/prometheus"
	"path/filepath"
	"testing"
)

func New(
	t *testing.T,
	o ...Option,
) *Tester {
	t.Helper()
	result := &Tester{
		Reporter: memory.New(),
		Registry: prometheus.NewRegistry(),
		master: func() face.SaltClient {
			return salt.New("localhost", 1, "guard", "guard")
		},
		keySyncInterval: option.New().KeySyncInterval,
		litePath:        filepath.Join(t.TempDir(), "gosaltd.db"),
		t:               t,
	}

	for _, f := range o {
		f(result)
	}

	return result
}
