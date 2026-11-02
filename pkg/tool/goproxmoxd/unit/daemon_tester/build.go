package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/option"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func build(
	t *testing.T,
	o *option.Proxmox,
) *daemon.Daemon {
	t.Helper()
	result := daemon.New(
		o,
		logger.New(t.Context()),
		memory.New(),
		prometheus.NewRegistry(),
	)
	t.Cleanup(result.Close)

	return result
}
