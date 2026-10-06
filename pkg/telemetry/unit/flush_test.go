package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/telemetry"
	"github.com/funtimecoding/soil/pkg/telemetry/constant"
	"github.com/funtimecoding/soil/pkg/telemetry/record"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

func TestFlushWaitsForRecordDelivery(t *testing.T) {
	var received atomic.Int32
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				r *http.Request,
			) {
				received.Add(1)
				w.WriteHeader(http.StatusNoContent)
			},
		),
	)
	defer s.Close()
	host, port := upstream_tester.HostPort(t, s.URL)
	c := telemetry.New(host, port, true)
	c.Record(
		record.NewDomain(
			"alfa_list",
			constant.CommandLine,
			constant.User,
			constant.Success,
		),
	)
	c.Flush()
	assert.Integer(t, 1, int(received.Load()))
}
