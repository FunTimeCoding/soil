package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/daemon"
)

func (t *Tester) Daemon() *daemon.Daemon {
	t.t.Helper()

	if t.daemon != nil {
		return t.daemon
	}

	t.daemon = daemon.New(
		t.daemonOption(),
		logger.New(t.t.Context()),
		t.Reporter,
	)
	t.t.Cleanup(t.daemon.Close)

	return t.daemon
}
