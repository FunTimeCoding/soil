package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	proxmox "github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/funtimecoding/soil/pkg/proxmox/mock_client"
	ssh "github.com/funtimecoding/soil/pkg/ssh/face"
	sshMock "github.com/funtimecoding/soil/pkg/ssh/mock_client"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	gonetboxdMock "github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/option"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
	peer string,
) (*model_context_server.Server, func() error) {
	t.Helper()
	c := mock_client.New()
	t.Cleanup(c.Close)
	n := gonetboxdMock.New()
	t.Cleanup(n.Close)

	switch peer {
	case gonetboxdConstant.Identity.Name():
		n.Failure = outage_tester.RefusedDial()
	default:
		c.SetFailure(outage_tester.RefusedDial())
	}

	o := option.New()
	o.Inventory = inventory.New[instance.Instance](
		instance.Instance{Name: "alfa", Host: "alfa.example"},
	)
	o.Hypervisor = func(*instance.Instance) (proxmox.ProxmoxClient, error) {
		return c, nil
	}
	o.Shell = func(*instance.Instance) (ssh.SSHClient, error) {
		return sshMock.New(), nil
	}
	o.Netbox = func() netbox.Client { return n }
	d := daemon.New(o, logger.New(t.Context()), r, prometheus.NewRegistry())
	t.Cleanup(d.Close)
	result := model_context_server.New(
		t,
		func(g *guard.Mux) { d.Mount(r, mock_recorder.New(), g) },
	)
	t.Cleanup(result.Close)

	return result, outage_tester.Guard(t, d.Workers(), peer).Unavailable
}
