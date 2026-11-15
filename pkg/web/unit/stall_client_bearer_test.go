package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/log/log_tester"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"testing"
)

func TestStallClientBearerAuthReachesBothViews(t *testing.T) {
	s := upstream_server.New(t)
	p := stall_client.New().BearerAuth("alfa").Build()
	defer p.Close()
	web.GetString(p.Client(), s.URL)
	assert.String(t, "Bearer alfa", s.Last().Header.Get(constant.Authorization))
	web.GetString(p.Halting(), s.URL)
	assert.Integer(t, 2, s.Count())
	assert.String(t, "Bearer alfa", s.Last().Header.Get(constant.Authorization))
}

func TestStallClientVerboseNeverLogsBearer(t *testing.T) {
	s := upstream_server.New(t)
	b := log_tester.Capture(t)
	p := stall_client.New().Verbose().BearerAuth("alfa").Build()
	defer p.Close()
	web.GetString(p.Client(), s.URL)
	assert.String(t, "Bearer alfa", s.Last().Header.Get(constant.Authorization))
	assert.StringContains(t, "REQUEST:", b.String())
	assert.StringNotContains(t, "alfa", b.String())
}
