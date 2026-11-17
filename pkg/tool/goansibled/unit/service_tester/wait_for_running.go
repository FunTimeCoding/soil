package service_tester

import (
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/time/wait_tester"
	"time"
)

func (o *Tester) WaitForRunning() {
	o.t.Helper()
	wait_tester.Until(
		o.t,
		func() bool {
			latest, found := o.latest()

			return found && latest.Status == provisionConstant.StoreStatusRunning
		},
		30*time.Second,
	)
}
