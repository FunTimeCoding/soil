package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"path/filepath"
	"testing"
	"time"
)

func New(
	t *testing.T,
	o ...Option,
) *Tester {
	t.Helper()
	result := &Tester{
		Path:         filepath.Join(t.TempDir(), "goalertlogd.db"),
		Alertmanager: mock_client.New(t),
		Reporter:     memory.New(),
		clock:        time.Now,
		retention:    constant.Retention,
		notifier:     notifier.New(),
		t:            t,
	}

	for _, f := range o {
		f(result)
	}

	return result
}
