package model_context_tester

import "github.com/funtimecoding/soil/pkg/tool/goalertlogd/daemon"

func (o *Tester) Daemon() *daemon.Daemon {
	return o.server.Daemon
}
