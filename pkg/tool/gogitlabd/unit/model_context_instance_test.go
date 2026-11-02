package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/daemon_tester"
	"testing"
)

func TestModelContextWorksOnTheSelectedInstance(t *testing.T) {
	r := memory.New()
	d, clients := daemon_tester.New(
		t,
		r,
		constant.FixtureInstance,
		constant.FixtureOtherInstance,
	)
	clients[constant.FixtureInstance].SeedProject(1, "group", "app")
	clients[constant.FixtureOtherInstance].SeedProject(2, "group", "web")
	v := startServer(t, d, r)
	defer v.Close()
	c := v.Connect(t)
	defer c.Close()
	assert.StringContains(
		t,
		"no instance selected",
		c.MustCallToolError(constant.ListProjects, nil),
	)
	c.MustCallTool(
		inventoryConstant.UseInstance,
		map[string]any{inventoryConstant.InstanceParameter: "bravo"},
	)
	assert.StringContains(
		t,
		"\"web\"",
		c.MustCallTool(constant.ListProjects, nil),
	)
}
