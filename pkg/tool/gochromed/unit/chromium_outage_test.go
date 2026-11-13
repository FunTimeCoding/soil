package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/unit/daemon_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestToolsOnChromiumOutage(t *testing.T) {
	outage_tester.Tools(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outageServer(t, r)
		},
		webConstant.OutageReaches,
		nil,
		nil,
	)
}

func TestDaemonHandsOutChromiumGuard(t *testing.T) {
	_, okay := daemon_tester.New(t, mock_client.New()).Workers()[0].(*peer.Guard)
	assert.True(t, okay)
}

func TestDaemonClosesClientItBuilt(t *testing.T) {
	c := mock_client.New()
	daemon_tester.New(t, c).Close()
	assert.Integer(t, 1, c.Closed())
}
