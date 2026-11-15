package outage_server

import (
	bandConstant "github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/band/face"
	"github.com/funtimecoding/soil/pkg/band/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/option"
	gonetboxdMock "github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
	peer string,
) (*model_context_server.Server, func() error) {
	t.Helper()
	engine := mock_client.New()
	t.Cleanup(engine.Close)
	engine.SetPower(
		bandConstant.PowerOn,
		bandConstant.PowerOn,
		bandConstant.PowerOffSoftGraceful,
	)
	n := gonetboxdMock.New()
	t.Cleanup(n.Close)

	switch peer {
	case gonetboxdConstant.Identity.Name():
		n.Failure = outage_tester.RefusedDial()
	default:
		engine.Fail(outage_tester.RefusedDial())
	}

	o := option.New()
	o.Inventory = inventory.New[instance.Instance](
		instance.Instance{Name: "alfa", Host: "alfa.example"},
	)
	o.Engine = func(*instance.Instance) (face.BandClient, error) {
		return engine, nil
	}
	o.Netbox = func() netbox.Client { return n }
	d := daemon.New(o, logger.New(t.Context()), r)
	t.Cleanup(d.Close)
	result := model_context_server.New(
		t,
		func(g *guard.Mux) { d.Mount(r, mock_recorder.New(), g) },
	)
	t.Cleanup(result.Close)

	return result, outage_tester.Guard(t, d.Workers(), peer).Unavailable
}
