package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/daemon_tester"
	"testing"
)

func TestModelContextAnswersClassifiedRefusalQuietly(t *testing.T) {
	r := memory.New()
	d, clients := daemon_tester.New(t, r, constant.FixtureInstance)
	clients[constant.FixtureInstance].FailProjects(
		not_found.New("group", "kilo"),
	)
	v := startServer(t, d, r)
	defer v.Close()
	c := v.Connect(t)
	defer c.Close()
	assert.StringContains(
		t,
		"group not found: kilo",
		c.MustCallToolError(constant.ListProjects, nil),
	)
	assert.Count(t, 0, r.Events())
}

func TestModelContextCapturesUnclassifiedFailure(t *testing.T) {
	r := memory.New()
	d, clients := daemon_tester.New(t, r, constant.FixtureInstance)
	clients[constant.FixtureInstance].FailProjects(
		errors.New("refused for no known reason"),
	)
	v := startServer(t, d, r)
	defer v.Close()
	c := v.Connect(t)
	defer c.Close()
	assert.StringContains(
		t,
		"refused for no known reason",
		c.MustCallToolError(constant.ListProjects, nil),
	)
	assert.Count(t, 1, r.Events())
}
