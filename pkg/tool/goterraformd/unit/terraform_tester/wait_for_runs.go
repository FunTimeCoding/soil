package terraform_tester

import (
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/time/wait_tester"
	"time"
)

func (o *Tester) WaitForRuns(count int) {
	o.t.Helper()
	wait_tester.Until(
		o.t,
		func() bool {
			runs := o.runs()

			if len(runs) < count {
				return false
			}

			for _, r := range runs {
				if r.Status == provisionConstant.StoreStatusRunning {
					return false
				}
			}

			return true
		},
		10*time.Second,
	)
}
