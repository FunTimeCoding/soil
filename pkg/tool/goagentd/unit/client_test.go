package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/system/run/stream/mock_stream"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/client"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/client/mock_client"
	agent "github.com/funtimecoding/soil/pkg/tool/goagentd/face"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestTheClientRunsAndReadsTheAgent(t *testing.T) {
	s := mock_stream.New([]string{"lima"}, true, nil)
	defer s.Close()
	d := daemon_tester.New(t, service_tester.Fixed(s))
	c := mountedClient(t, d, constant.ModelContextTestToken)
	assert.FatalOnError(t, c.Reachable(t.Context()))
	assert.True(t, validation.Is(c.PostRun(t.Context(), "")))
	assert.FatalOnError(t, c.PostRun(t.Context(), "mike"))
	assert.True(t, conflict.Is(c.PostRun(t.Context(), "november")))
	running, e := c.GetStatus(t.Context())
	assert.FatalOnError(t, e)
	assert.String(t, "running", running.State)
	d.Workers()[0].(face.Drainer).Interrupt()
	done := awaitServedDone(t, c)
	assert.String(t, "done", done.State)
	assert.String(t, "lima", *done.Result)
}

func TestTheClientHearsAClosingRunner(t *testing.T) {
	d := daemon_tester.New(
		t,
		service_tester.Fixed(mock_stream.New(nil, false, nil)),
	)
	c := mountedClient(t, d, constant.ModelContextTestToken)
	d.Close()
	assert.True(t, unavailable.Is(c.PostRun(t.Context(), "oscar")))
}

func TestARefusedTokenIsAnError(t *testing.T) {
	d := daemon_tester.New(
		t,
		service_tester.Fixed(mock_stream.New(nil, false, nil)),
	)
	c := mountedClient(t, d, "papa-wrong")
	_, e := c.GetStatus(t.Context())
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "status status: 401", e.Error())
	assert.NotNil(t, c.Reachable(t.Context()))
}

func TestTheClientAndItsMockKeepTheFace(t *testing.T) {
	r := client.New(locator.New("localhost").Insecure(), "quebec")
	m := mock_client.New()
	clients := []agent.Client{r, m}

	for _, c := range clients {
		c.Close()
	}

	assert.Integer(t, 1, m.Closed)
}
