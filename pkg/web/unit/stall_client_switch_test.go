package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/log/log_tester"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"testing"
)

func TestStallClientBasicAuthReachesBothViews(t *testing.T) {
	s := upstream_server.New(t)
	p := stall_client.New().BasicAuth("alfa", "bravo").Build()
	defer p.Close()
	web.GetString(p.Client(), s.URL)
	user, password, _ := s.Last().BasicAuth()
	assert.String(t, "alfa", user)
	assert.String(t, "bravo", password)
	web.GetString(p.Halting(), s.URL)
	assert.Integer(t, 2, s.Count())
	user, password, _ = s.Last().BasicAuth()
	assert.String(t, "alfa", user)
	assert.String(t, "bravo", password)
}

func TestStallClientVerboseNeverLogsCredentials(t *testing.T) {
	s := upstream_server.New(t)
	b := log_tester.Capture(t)
	p := stall_client.New().Verbose().BasicAuth("alfa", "bravo").Build()
	defer p.Close()
	web.GetString(p.Client(), s.URL)
	user, password, _ := s.Last().BasicAuth()
	assert.String(t, "alfa", user)
	assert.String(t, "bravo", password)
	assert.StringContains(t, "REQUEST:", b.String())
	assert.StringNotContains(t, "Authorization", b.String())
}
