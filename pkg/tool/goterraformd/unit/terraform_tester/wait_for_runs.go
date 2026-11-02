package terraform_tester

import (
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"time"
)

func (o *Tester) WaitForRuns(count int) {
	o.t.Helper()
	deadline := time.Now().Add(10 * time.Second)

	for time.Now().Before(deadline) {
		runs := o.runs()
		finished := len(runs) >= count

		for _, r := range runs {
			if r.Status == provisionConstant.StoreStatusRunning {
				finished = false
			}
		}

		if finished {
			return
		}

		time.Sleep(10 * time.Millisecond)
	}

	o.t.Fatalf("timed out waiting for %d finished runs", count)
}
