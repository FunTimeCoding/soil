package service_tester

import (
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
	deadline := time.Now().Add(5 * time.Second)

	for time.Now().Before(deadline) {
		if s := v.Status(); s.State == constant.RunnerDone {
			return s
		}

		time.Sleep(10 * time.Millisecond)
	}

	t.Fatalf("run did not finish")

	return nil
}
