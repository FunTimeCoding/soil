package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	"testing"
)

func TestModelContextWorksOnTheSelectedInstance(t *testing.T) {
	alfa := mock_client.New()
	alfa.SeedProject(1, "group", "app")
	bravo := mock_client.New()
	bravo.SeedProject(2, "group", "web")
	v := startServer(
		t,
		service_tester.Service(
			map[string]*mock_client.Client{
				constant.FixtureInstance:      alfa,
				constant.FixtureOtherInstance: bravo,
			},
		),
	)
	defer v.Stop()
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
