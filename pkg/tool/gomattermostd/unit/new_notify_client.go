package unit

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/connector"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newNotifyClient(t *testing.T) (*notifySink, *connector.Client) {
	t.Helper()
	result := &notifySink{}
	server := httptest.NewServer(
		http.HandlerFunc(
			func(
				_ http.ResponseWriter,
				q *http.Request,
			) {
				var v client.NotifyRequest
				notation.MustDecode(string(system.ReadAll(q.Body)), &v, false)
				result.add(v)
			},
		),
	)
	t.Cleanup(server.Close)

	return result, connector.New(server.URL, false, "")
}
