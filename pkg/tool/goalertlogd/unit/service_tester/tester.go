package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/service"
	"testing"
	"time"
)

type Tester struct {
	Path         string
	Alertmanager *mock_client.Client
	Reporter     *memory.Memory
	clock        func() time.Time
	retention    time.Duration
	notifier     face.EventNotifier
	service      *service.Service
	daemon       *daemon.Daemon
	t            *testing.T
}
