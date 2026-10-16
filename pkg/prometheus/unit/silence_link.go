package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	"net"
	"net/http/httptest"
	"testing"
)

func silenceLink(
	t *testing.T,
	s *httptest.Server,
	o ...alertmanager.Option,
) string {
	t.Helper()
	p := prometheus.New("127.0.0.1", 0, false, "", "")
	t.Cleanup(p.Close)
	c := alertmanager.New(
		"127.0.0.1",
		s.Listener.Addr().(*net.TCPAddr).Port,
		false,
		"",
		"",
		p,
		o...,
	)
	t.Cleanup(c.Close)
	v, e := c.Silences(true)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)

	return v[0].Link
}
