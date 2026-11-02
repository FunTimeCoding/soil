package service_tester

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/service"
	"testing"
)

func NewService(
	t *testing.T,
	path string,
	a face.AlertmanagerClient,
) *service.Service {
	t.Helper()

	return NewServiceWith(t, path, a, constant.Retention, notifier.New())
}
