package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newRefusingAlertmanager(
	t *testing.T,
	status int,
	body string,
) *alertmanager.Client {
	t.Helper()
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				w.Header().Set(webConstant.ContentType, webConstant.Object)
				w.WriteHeader(status)
				_, e := w.Write([]byte(body))
				errors.PanicOnError(e)
			},
		),
	)
	t.Cleanup(s.Close)
	result := alertmanager.New(
		"127.0.0.1",
		s.Listener.Addr().(*net.TCPAddr).Port,
		false,
		"",
		"",
	)
	t.Cleanup(result.Close)

	return result
}
