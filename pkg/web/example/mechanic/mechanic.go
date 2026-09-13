package mechanic

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Mechanic(
	address string,
	extended string,
	serverSide string,
) {
	r := memory.New()
	u := New(notifier.New(), extended, serverSide)
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			server.New(
				identity.Example(),
				address,
				func(m *http.ServeMux) {
					u.Mount(guard.New(m, []string{constant.MechanicToken}))
				},
			).WithMiddleware(u.Recovery(r)),
		),
	).RunUntilSignal()
}
