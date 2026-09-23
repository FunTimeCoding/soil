package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/connector"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewSink(t *testing.T) (*Sink, *connector.Client) {
	t.Helper()
	result := &Sink{}
	server := httptest.NewServer(
		http.HandlerFunc(
			func(
				_ http.ResponseWriter,
				q *http.Request,
			) {
				var v client.NotifyRequest
				notation.MustDecode(string(system.ReadAll(q.Body)), &v, false)
				result.Add(v)
			},
		),
	)
	t.Cleanup(server.Close)

	return result, connector.New(server.URL, false, "")
}
