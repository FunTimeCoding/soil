package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newRefusingClient(
	t *testing.T,
	status int,
	body string,
) *prometheus.Client {
	t.Helper()
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				r *http.Request,
			) {
				w.Header().Set("Content-Type", "application/json")

				if r.Method == http.MethodGet {
					_, e := w.Write(
						[]byte(
							`{"status":"success","data":{"resultType":"vector","result":[]}}`,
						),
					)
					errors.PanicOnError(e)

					return
				}

				w.WriteHeader(status)
				_, e := w.Write([]byte(body))
				errors.PanicOnError(e)
			},
		),
	)
	t.Cleanup(s.Close)
	result := prometheus.New(
		"127.0.0.1",
		s.Listener.Addr().(*net.TCPAddr).Port,
		false,
		"",
		"",
	)
	t.Cleanup(result.Close)

	return result
}
