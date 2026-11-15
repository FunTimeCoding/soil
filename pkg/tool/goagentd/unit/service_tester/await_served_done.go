package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/time/wait_tester"
	goagentd "github.com/funtimecoding/soil/pkg/tool/goagentd/client"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/generated/client"
	"testing"
	"time"
)

func AwaitServedDone(
	t *testing.T,
	c *goagentd.Client,
) client.StatusResponse {
	t.Helper()
	var result client.StatusResponse
	wait_tester.Until(
		t,
		func() bool {
			s, e := c.GetStatus(t.Context())
			assert.FatalOnError(t, e)
			result = s

			return s.State == string(constant.RunnerDone)
		},
		5*time.Second,
	)

	return result
}
