package unit_test

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	lifecycleServer "github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/system"
	generated "github.com/funtimecoding/soil/pkg/tool/goraidparsed/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/server"
	"net/http"
	"testing"
)

func TestRunLifecycle(t *testing.T) {
	port, n := system.ClaimPort()
	l := lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			lifecycleServer.New(
				"",
				func(m *http.ServeMux) {
					generated.HandlerFromMux(
						generated.NewStrictHandler(
							server.New("", "", t.TempDir(), nil, nil),
							nil,
						),
						m,
					)
				},
			).WithListener(n),
		),
	)
	l.Run()
	defer l.Stop()
	assert.Listen(t, port)
	base := fmt.Sprintf("http://localhost:%d", port)
	assert.HTTPStatus(
		t,
		fmt.Sprintf("%s/api/v1/status", base),
		http.StatusOK,
	)
}
