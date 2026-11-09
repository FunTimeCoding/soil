package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/daemon"
	"testing"
	"time"
)

func New(
	t *testing.T,
	clients map[string]*mock_client.Client,
) *daemon.Daemon {
	t.Helper()

	return NewWith(t, clients, time.Now)
}
