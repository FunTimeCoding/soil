package service_tester

import (
	"github.com/funtimecoding/soil/pkg/time/wait_tester"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/types/run_result"
	"time"
)

func (o *Tester) AwaitDone() *run_result.Result {
	o.t.Helper()
	var result *run_result.Result
	wait_tester.Until(
		o.t,
		func() bool {
			result = o.Service().Status()

			return result.State == constant.RunnerDone
		},
		5*time.Second,
	)

	return result
}
