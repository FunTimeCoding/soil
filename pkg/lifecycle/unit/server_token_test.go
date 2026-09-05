package unit

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	strings "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"testing"
)

func TestRunServerTokens(t *testing.T) {
	p, n := system.ClaimPort()
	l := lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			server.New(
				identity.Example(),
				"",
				func(m *http.ServeMux) {
					m.HandleFunc(
						fmt.Sprintf("GET %s", webConstant.EchoPath),
						func(
							w http.ResponseWriter,
							_ *http.Request,
						) {
							w.WriteHeader(http.StatusOK)
						},
					)
				},
			).WithListener(n).WithTokens(
				[]string{strings.UpperAlfa},
			),
		),
	)
	l.Run()
	defer l.Stop()
	assert.Listen(t, p)
	assert.HTTPStatus(
		t,
		fmt.Sprintf("http://localhost:%d/health", p),
		http.StatusOK,
	)
	assert.HTTPStatus(
		t,
		fmt.Sprintf("http://localhost:%d/version", p),
		http.StatusOK,
	)
	assert.HTTPStatus(
		t,
		fmt.Sprintf("http://localhost:%d/echo", p),
		http.StatusUnauthorized,
	)
	q := web.NewGet(fmt.Sprintf("http://localhost:%d/echo", p))
	web.Bearer(q, strings.UpperAlfa)
	r, e := web.Client().Do(q)
	assert.FatalOnError(t, e)
	defer errors.PanicClose(r.Body)
	assert.Integer(t, http.StatusOK, r.StatusCode)
}
