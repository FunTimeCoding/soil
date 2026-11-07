package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/proxmox/mock_client"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	c := mock_client.New()
	c.AddNode("test")
	d := daemon_tester.NewWithNetbox(
		t,
		map[string]*mock_client.Client{"test": c},
		netbox.New(),
	)
	v := model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(memory.New(), mock_recorder.New(), g)
		},
	)
	result := &Server{MockClient: c, Server: v}
	t.Cleanup(result.Close)

	return result
}
