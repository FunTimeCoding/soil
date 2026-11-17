package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	"testing"
)

func TestModelContextWorksOnSelectedInstance(t *testing.T) {
	s := base.New(
		t,
		service_tester.WithInstances(
			constant.FixtureInstance,
			constant.FixtureOtherInstance,
		),
	)
	s.Alfa().SeedProject(1, "group", "app")
	s.Bravo().SeedProject(2, "group", "web")
	c := s.Tools()
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
