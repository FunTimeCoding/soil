package unit

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	lifecycleServer "github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/system"
	generated "github.com/funtimecoding/soil/pkg/tool/goraidparsed/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/server"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"net/http"
	"testing"
)

func TestRunLifecycle(t *testing.T) {
	port, n := system.ClaimPort()
	l := lifecycle.New(logger.New(context.Background())).WithServer(
		lifecycleServer.New(
			identity.Example(),
			"",
			func(m *http.ServeMux) {
				generated.HandlerFromMux(
					generated.NewStrictHandler(
						server.New(service_tester.New(t), memory.New()),
						nil,
					),
					m,
				)
			},
		).WithListener(n),
	)
	l.Run()
	defer l.Close()
	p := stall_client.New().Build()
	defer p.Close()
	c := p.Client()
	assert.Listen(t, port)
	assert.HTTPStatus(
		t,
		c,
		fmt.Sprintf("http://localhost:%d/api/status", port),
		http.StatusOK,
	)
}
