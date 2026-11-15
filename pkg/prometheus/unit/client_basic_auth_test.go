package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"testing"
)

func TestClientSendsBasicAuth(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(`{"status":"success","data":["job"]}`),
	)
	host, port := upstream_server.HostPort(t, s.URL)
	c := prometheus.New(host, port, false, "alfa", "bravo")
	defer c.Close()
	labels, e := c.AllLabels()
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"job"}, labels)
	user, password, _ := s.Last().BasicAuth()
	assert.String(t, "alfa", user)
	assert.String(t, "bravo", password)
}
