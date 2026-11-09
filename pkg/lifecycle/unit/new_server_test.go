package unit

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/option"
	"github.com/funtimecoding/soil/pkg/web/route"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"net/http"
	"testing"
)

func TestServerTakesItsAddressAndTokensFromWebOption(t *testing.T) {
	p := system.FindUnusedPort(19400)
	l := lifecycle.New(logger.New(context.Background())).WithServer(
		server.New(
			identity.Example(),
			option.New(fmt.Sprintf("127.0.0.1:%d", p), "token"),
			memory.New(),
			func(g *guard.Mux) {
				g.Token(
					route.Get("/alfa"),
					func(w http.ResponseWriter, _ *http.Request) {
						w.WriteHeader(http.StatusNoContent)
					},
				)
			},
		),
	)
	l.Run()
	defer l.Close()
	o := stall_client.New().Build()
	defer o.Close()
	assert.Listen(t, p)
	locator := fmt.Sprintf("http://localhost:%d/alfa", p)
	bare, e := o.Client().Do(web.NewGet(locator))
	assert.FatalOnError(t, e)
	defer errors.PanicClose(bare.Body)
	assert.Integer(t, http.StatusUnauthorized, bare.StatusCode)
	q := web.NewGet(locator)
	web.Bearer(q, "token")
	signed, f := o.Client().Do(q)
	assert.FatalOnError(t, f)
	defer errors.PanicClose(signed.Body)
	assert.Integer(t, http.StatusNoContent, signed.StatusCode)
}
