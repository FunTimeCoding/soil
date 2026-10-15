package ansible_tester

import (
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/model/run"
	"time"
)

func (o *Tester) WaitForRun() *run.Run {
	o.t.Helper()
	deadline := time.Now().Add(30 * time.Second)

	for time.Now().Before(deadline) {
		runs, e := o.Store.Recent(1)

		if e == nil && len(runs) == 1 &&
			runs[0].Status != provisionConstant.StoreStatusRunning {
			return &runs[0]
		}

		time.Sleep(50 * time.Millisecond)
	}

	o.t.Fatal("no finished run")

	return nil
}
