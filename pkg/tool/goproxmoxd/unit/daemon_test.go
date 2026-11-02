package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/proxmox/mock_client"
	sshMock "github.com/funtimecoding/soil/pkg/ssh/mock_client"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"go.uber.org/goleak"
	"net/http"
	"testing"
)

func TestTheDaemonHandsOneGuardPerHypervisorOutBeforeThePoller(t *testing.T) {
	alfa := mock_client.New()
	defer alfa.Close()
	bravo := mock_client.New()
	defer bravo.Close()
	d := daemon_tester.New(
		t,
		map[string]*mock_client.Client{"alfa": alfa, "bravo": bravo},
	)
	workers := d.Workers()
	assert.Count(t, 3, workers)
	_, first := workers[0].(*peer.Guard)
	_, second := workers[1].(*peer.Guard)
	_, third := workers[2].(*peer.Guard)
	assert.True(t, first)
	assert.True(t, second)
	assert.False(t, third)
}

func TestTheDaemonHandsEveryGuardOutBeforeItsClocks(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	d := daemon_tester.NewWithNetbox(
		t,
		map[string]*mock_client.Client{"alfa": c},
		netbox.New(),
	)
	workers := d.Workers()
	assert.Count(t, 4, workers)
	_, first := workers[0].(*peer.Guard)
	_, second := workers[1].(*peer.Guard)
	_, third := workers[2].(*peer.Guard)
	assert.True(t, first)
	assert.True(t, second)
	assert.False(t, third)
}

func TestClosingTheDaemonClosesEveryClientItBuilt(t *testing.T) {
	alfa := mock_client.New()
	defer alfa.Close()
	bravo := mock_client.New()
	defer bravo.Close()
	n := netbox.New()
	defer n.Close()
	d := daemon_tester.NewWithNetbox(
		t,
		map[string]*mock_client.Client{"alfa": alfa, "bravo": bravo},
		n,
	)
	d.Close()
	assert.Integer(t, 1, alfa.Closed)
	assert.Integer(t, 1, bravo.Closed)
	assert.Integer(t, 1, n.Closed)
}

func TestClosingTheDaemonClosesTheShellsItDialled(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	h := sshMock.New()
	defer h.Close()
	d := daemon_tester.NewWithShells(
		t,
		map[string]*mock_client.Client{"alfa": c},
		map[string]*sshMock.Client{"alfa": h},
	)
	v := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			d.Mount(memory.New(), mock_recorder.New(), g)
		},
	)
	defer v.Close()
	m := v.Connect(t)
	defer m.Close()
	m.MustCallTool(constant.ListSnippets, nil)
	d.Close()
	assert.Integer(t, 1, h.Closed)
}

func TestTheDaemonLeavesNothingRunning(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	before := goleak.IgnoreCurrent()
	d := daemon_tester.NewWithNetbox(
		t,
		map[string]*mock_client.Client{"alfa": c},
		netbox.New(),
	)

	for _, w := range d.Workers() {
		w.Start()
	}

	d.Close()
	goleak.VerifyNone(t, before)
}
