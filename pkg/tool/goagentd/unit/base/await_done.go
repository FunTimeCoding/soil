package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/time/wait_tester"
	goagentd "github.com/funtimecoding/soil/pkg/tool/goagentd/client"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/generated/client"
	"time"
)

func (s *Server) AwaitDone(c *goagentd.Client) client.StatusResponse {
	s.t.Helper()
	var result client.StatusResponse
	wait_tester.Until(
		s.t,
		func() bool {
			r, e := c.GetStatus(s.t.Context())
			assert.FatalOnError(s.t, e)
			result = r

			return r.State == string(constant.RunnerDone)
		},
		5*time.Second,
	)

	return result
}
