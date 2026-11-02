package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/client"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/constant"
	generated "github.com/funtimecoding/soil/pkg/tool/goagentd/generated/client"
	"testing"
	"time"
)

func awaitServedDone(
	t *testing.T,
	c *client.Client,
) generated.StatusResponse {
	t.Helper()
	deadline := time.Now().Add(5 * time.Second)

	for time.Now().Before(deadline) {
		s, e := c.GetStatus(t.Context())
		assert.FatalOnError(t, e)

		if s.State == string(constant.RunnerDone) {
			return s
		}

		time.Sleep(10 * time.Millisecond)
	}

	t.Fatalf("run did not finish")

	return generated.StatusResponse{}
}
