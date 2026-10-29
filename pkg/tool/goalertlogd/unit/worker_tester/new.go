package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/service_tester"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()

	return NewWith(
		t,
		service_tester.DatabasePath(t),
		service_tester.NewAlertmanager(t),
		constant.Retention,
	)
}
