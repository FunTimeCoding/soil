//go:build browser

package mechanic

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/example/mechanic"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func serve(t *testing.T) string {
	t.Helper()
	p, n := system.ClaimPort()
	u := mechanic.New(notifier.New(), extendedSource(), serverSideSource())
	b := lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			server.New(
				identity.Example(),
				"",
				func(m *http.ServeMux) {
					u.Mount(guard.New(m, []string{constant.MechanicToken}))
				},
			).WithListener(n).WithMiddleware(u.Recovery(memory.New())),
		),
	)
	b.Run()
	t.Cleanup(b.Stop)
	assert.Listen(t, p)

	return fmt.Sprintf("http://localhost:%d", p)
}
