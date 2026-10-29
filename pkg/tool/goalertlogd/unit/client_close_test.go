package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/client"
	generated "github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

func TestClientCloseDropsIdleConnections(t *testing.T) {
	var opened atomic.Int64
	s := httptest.NewUnstartedServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				w.Header().Set(webConstant.ContentType, webConstant.Object)
				_, e := w.Write([]byte("[]"))
				errors.PanicOnError(e)
			},
		),
	)
	s.Config.ConnState = func(
		_ net.Conn,
		c http.ConnState,
	) {
		if c == http.StateNew {
			opened.Add(1)
		}
	}
	s.Start()
	defer s.Close()
	c := client.New(
		locator.New("127.0.0.1").Port(
			s.Listener.Addr().(*net.TCPAddr).Port,
		).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	p := &generated.GetAlertsParams{Name: "HighMemory"}
	_, f := c.GetAlerts(t.Context(), p)
	assert.FatalOnError(t, f)
	c.Close()
	_, g := c.GetAlerts(t.Context(), p)
	assert.FatalOnError(t, g)
	assert.Integer(t, int64(2), opened.Load())
}
