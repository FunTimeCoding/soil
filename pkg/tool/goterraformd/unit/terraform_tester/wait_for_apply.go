package terraform_tester

import (
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"time"
)

func (o *Tester) waitForApply() {
	o.t.Helper()
	deadline := time.Now().Add(10 * time.Second)

	for time.Now().Before(deadline) {
		runs, e := o.Service.Recent(1)

		if e == nil &&
			len(runs) == 1 &&
			runs[0].Status != provisionConstant.StoreStatusRunning {
			return
		}

		time.Sleep(10 * time.Millisecond)
	}

	o.t.Fatal("timed out waiting for the startup apply")
}
