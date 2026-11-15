package service_tester

import (
	"github.com/funtimecoding/soil/pkg/time/wait_tester"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/service"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/types/run_result"
	"testing"
	"time"
)

func AwaitDone(
	t *testing.T,
	v *service.Service,
) *run_result.Result {
	t.Helper()
	var result *run_result.Result
	wait_tester.Until(
		t,
		func() bool {
			result = v.Status()

			return result.State == constant.RunnerDone
		},
		5*time.Second,
	)

	return result
}
