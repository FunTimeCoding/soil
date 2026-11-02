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
		latest, found := o.latest()

		if found && latest.Status != provisionConstant.StoreStatusRunning {
			return latest
		}

		time.Sleep(50 * time.Millisecond)
	}

	o.t.Fatal("no finished run")

	return nil
}
