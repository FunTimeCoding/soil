package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/opnsense/unit/opnsense_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestRestartingOpnsenseIsUnavailable(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithTLS(),
		upstream_server.WithStatus(http.StatusServiceUnavailable),
	)
	_, e := opnsense_tester.New(t, s.URL).Leases("")
	assert.True(t, unavailable.Is(e))
}
