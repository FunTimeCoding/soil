package outage_tester

import (
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
)

func (o *Tester) Unavailable() error {
	o.t.Helper()

	return refusal.New(
		"list issues",
		upstream_server.Answered(o.t, http.StatusServiceUnavailable),
		nil,
	)
}
