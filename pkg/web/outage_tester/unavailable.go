package outage_tester

import (
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func Unavailable(t *testing.T) error {
	t.Helper()

	return refusal.New(
		"list issues",
		upstream_server.Answered(t, http.StatusServiceUnavailable),
		nil,
	)
}
