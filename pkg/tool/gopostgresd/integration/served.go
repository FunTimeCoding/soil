package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/dump"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/option"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func served(
	t *testing.T,
	container string,
	i *instance.Instance,
	directory string,
) (*client.ClientWithResponses, *memory.Memory) {
	t.Helper()
	o := option.New()
	o.Inventory = inventory.New[instance.Instance](*i)
	o.BackupDirectory = directory
	o.Dump = dump.New(
		"podman",
		"exec",
		"--env",
		constant.PasswordEnvironment,
		"--env",
		constant.ApplicationNameEnvironment,
		container,
	)
	d := daemon.New(o, logger.New(t.Context()), memory.New())
	t.Cleanup(d.Close)
	r := memory.New()
	s := model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(r, mock_recorder.New(), g)
		},
	)
	t.Cleanup(s.Close)
	result, e := client.NewClientWithResponses(
		s.Address(webConstant.RootPath),
		client.WithHTTPClient(s.Client()),
		client.WithRequestEditorFn(
			web.BearerEditor(generativeConstant.ModelContextTestToken),
		),
	)
	assert.FatalOnError(t, e)

	return result, r
}
