package ansible_tester

import (
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"time"
)

func (o *Tester) WaitForRunning() {
	o.t.Helper()
	deadline := time.Now().Add(30 * time.Second)

	for time.Now().Before(deadline) {
		runs, e := o.Service.Recent(1)

		if e == nil && len(runs) == 1 &&
			runs[0].Status == provisionConstant.StoreStatusRunning {
			return
		}

		time.Sleep(50 * time.Millisecond)
	}

	o.t.Fatal("no run in flight")
}
