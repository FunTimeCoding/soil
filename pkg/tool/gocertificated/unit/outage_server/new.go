package outage_server

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/server_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/service_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
) (*model_context_server.Server, func() error) {
	t.Helper()
	s := base.New(t, service_tester.WithReporter(r))
	s.Forge.Fail(outage_tester.RefusedDial())
	_, e := s.GeneratedClient().PostAuthorityWithBodyWithResponse(
		t.Context(),
		webConstant.Object,
		bytes.NewReader(notation.Marshal(server_tester.RootBody())),
	)
	assert.FatalOnError(t, e)

	return s.Served(), outage_tester.New(t).Guard(
		s.Daemon().Workers(),
		constant.ForgePeer,
	).Unavailable
}
