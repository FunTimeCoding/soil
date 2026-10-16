package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAlertmanagerClientDialsConfiguredPort(t *testing.T) {
	var seen string
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				r *http.Request,
			) {
				seen = r.URL.Path
				w.Header().Set(webConstant.ContentType, webConstant.Object)
				_, e := w.Write([]byte(`{}`))
				errors.PanicOnError(e)
			},
		),
	)
	defer s.Close()
	p := prometheus.New("127.0.0.1", 0, false, "", "")
	defer p.Close()
	c := alertmanager.New(
		"127.0.0.1",
		s.Listener.Addr().(*net.TCPAddr).Port,
		false,
		"",
		"",
		p,
	)
	defer c.Close()
	_, e := c.Status()
	assert.FatalOnError(t, e)
	assert.String(t, "/api/v2/status", seen)
}
