package postgres_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/dump"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/option"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
)

func (o *Tester) Serve(
	container string,
	i *instance.Instance,
	directory string,
) (*client.ClientWithResponses, *memory.Memory) {
	o.t.Helper()
	p := option.New()
	p.Inventory = inventory.New[instance.Instance](*i)
	p.BackupDirectory = directory
	p.Dump = dump.New(
		"podman",
		"exec",
		"--env",
		constant.PasswordEnvironment,
		"--env",
		constant.ApplicationNameEnvironment,
		container,
	)
	d := daemon.New(p, logger.New(o.t.Context()), memory.New())
	o.t.Cleanup(d.Close)
	r := memory.New()
	s := model_context_server.Mount(o.t, d, r)
	result, e := client.NewClientWithResponses(
		s.Address(webConstant.RootPath),
		client.WithHTTPClient(s.Client()),
		client.WithRequestEditorFn(
			web.BearerEditor(generativeConstant.ModelContextTestToken),
		),
	)
	assert.FatalOnError(o.t, e)

	return result, r
}
