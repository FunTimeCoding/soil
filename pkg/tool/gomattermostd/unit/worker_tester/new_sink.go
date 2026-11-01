package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/types/notify_sink"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewSink(t *testing.T) (*notify_sink.Sink, *goclauded.Client) {
	t.Helper()
	result := notify_sink.New()
	server := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				q *http.Request,
			) {
				var v client.NotifyRequest
				notation.MustDecode(string(system.ReadAll(q.Body)), &v, false)
				result.Add(v)
				web.EncodeNotation(w, client.DeliveryResponse{})
			},
		),
	)
	t.Cleanup(server.Close)
	c := goclauded.New(server.URL, false, "")
	t.Cleanup(c.Close)

	return result, c
}
