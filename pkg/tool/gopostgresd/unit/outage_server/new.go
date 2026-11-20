package outage_server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/store/backup"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
) (*model_context_server.Server, func() error) {
	t.Helper()
	s := base.New(t, service_tester.WithReporter(r))
	s.Alfa().Failure = outage_tester.RefusedDial()
	path, e := backup.New(s.BackupDirectory).Prepare(
		"alfa",
		"2030-01-02T03-04-05",
	)
	assert.FatalOnError(t, e)
	system.SaveFile(path, "dump")

	return s.Served(), outage_tester.New(t).Guard(
		s.Daemon().Workers(),
		join.Colon(constant.InstancePeer, "alfa"),
	).Unavailable
}
