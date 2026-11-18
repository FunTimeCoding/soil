package battery_tester

import "github.com/funtimecoding/soil/pkg/tool/gobatteryd/worker"

func (t *Tester) Poll() {
	w := t.Daemon.Workers()
	w[len(w)-1].(*worker.Worker).Poll()
}
