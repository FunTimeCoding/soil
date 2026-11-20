package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/option"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	goquerydConstant "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"path/filepath"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
) (*model_context_server.Server, func() error) {
	t.Helper()
	q := mock_client.New(t)
	q.Failure = outage_tester.RefusedDial()
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "memory.db")
	o.Query = func() face.Client { return q }
	o.ReferenceRoot = t.TempDir()
	d := daemon.New(o, logger.New(t.Context()), r)
	t.Cleanup(d.Close)
	result := model_context_server.Mount(t, d, r)

	return result, outage_tester.New(t).Guard(
		d.Workers(),
		goquerydConstant.Identity.Name(),
	).Unavailable
}
