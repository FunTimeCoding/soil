package worker_tester

import "github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"

func (o *Tester) Restart() *Tester {
	o.t.Helper()
	o.Worker.Close()

	return NewWith(o.t, o.Path, o.MockClient, constant.Retention)
}
