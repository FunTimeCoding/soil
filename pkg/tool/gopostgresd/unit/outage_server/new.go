package outage_server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/postgres/face"
	"github.com/funtimecoding/soil/pkg/postgres/mock_client"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/dump/mock_dump"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/option"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/store/backup"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
) (*model_context_server.Server, func() error) {
	t.Helper()
	c := mock_client.New()
	c.Failure = outage_tester.RefusedDial()
	o := option.New()
	o.Inventory = inventory.New[instance.Instance](
		instance.Instance{Name: "alfa", Host: "alfa.example", Database: "alfa"},
	)
	o.BackupDirectory = t.TempDir()
	o.Dump = mock_dump.New()
	o.Client = func(*instance.Instance) (face.PostgresClient, error) {
		return c, nil
	}
	path, e := backup.New(o.BackupDirectory).Prepare(
		"alfa",
		"2030-01-02T03-04-05",
	)
	assert.FatalOnError(t, e)
	system.SaveFile(path, "dump")
	d := daemon.New(o, logger.New(t.Context()), r)
	t.Cleanup(d.Close)
	result := model_context_server.New(
		t,
		func(g *guard.Mux) { d.Mount(r, mock_recorder.New(), g) },
	)
	t.Cleanup(result.Close)

	return result, outage_tester.Guard(
		t,
		d.Workers(),
		join.Colon(constant.InstancePeer, "alfa"),
	).Unavailable
}
