package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"slices"
	"strings"
	"testing"
)

func TestEveryToolFailsOnAChromiumOutage(t *testing.T) {
	r := memory.New()
	var unanswered []string

	for tool, text := range outage_tester.ToolFailures(
		t,
		func() *model_context_server.Server { return outageServer(t, r) },
		nil,
	) {
		if !strings.Contains(text, "connection refused") {
			unanswered = append(unanswered, tool)
		}
	}

	slices.Sort(unanswered)
	assert.String(t, "", strings.Join(unanswered, ", "))
	assert.Count(t, 0, r.Events())
}

func TestTheDaemonHandsOutTheChromiumGuard(t *testing.T) {
	_, okay := daemon_tester.New(t, mock_client.New()).Workers()[0].(*peer.Guard)
	assert.True(t, okay)
}

func TestTheDaemonClosesTheClientItBuilt(t *testing.T) {
	c := mock_client.New()
	daemon_tester.New(t, c).Close()
	assert.Integer(t, 1, c.Closed())
}
