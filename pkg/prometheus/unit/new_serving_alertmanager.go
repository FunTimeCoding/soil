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

func newServingAlertmanager(
	t *testing.T,
	payload string,
) (*alertmanager.Client, *[]string) {
	t.Helper()
	var queries []string
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				q *http.Request,
			) {
				queries = append(queries, q.URL.RawQuery)
				w.Header().Set(webConstant.ContentType, webConstant.Object)
				_, e := w.Write([]byte(payload))
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

	return result, &queries
}
