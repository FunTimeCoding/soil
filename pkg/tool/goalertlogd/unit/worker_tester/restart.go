package worker_tester

import "github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/service_tester"

func (o *Tester) Restart() *Tester {
	o.t.Helper()
	o.Worker.Close()

	return New(
		o.t,
		service_tester.WithPath(o.Path),
		service_tester.WithAlertmanager(o.Alertmanager),
	)
}
