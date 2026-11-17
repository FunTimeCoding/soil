package service_tester

import (
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/model/run"
	"github.com/funtimecoding/soil/pkg/time/wait_tester"
	"time"
)

func (o *Tester) WaitForRun() *run.Run {
	o.t.Helper()
	var result *run.Run
	wait_tester.Until(
		o.t,
		func() bool {
			latest, found := o.latest()
			result = latest

			return found && latest.Status != provisionConstant.StoreStatusRunning
		},
		30*time.Second,
	)

	return result
}
